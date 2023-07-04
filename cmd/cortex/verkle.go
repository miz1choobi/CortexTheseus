// Copyright 2022 The CortexTheseus Authors
// This file is part of CortexTheseus.
//
// CortexTheseus is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// CortexTheseus is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with CortexTheseus. If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"os"

	"github.com/CortexFoundation/CortexTheseus/cmd/utils"
	"github.com/CortexFoundation/CortexTheseus/common"
	"github.com/CortexFoundation/CortexTheseus/core/rawdb"
	"github.com/CortexFoundation/CortexTheseus/log"
	"github.com/gballet/go-verkle"
	//cli "github.com/urfave/cli/v2"
	cli "gopkg.in/urfave/cli.v1"
)

var (
	zero          [32]byte
	verkleCommand = cli.Command{
		Name:        "verkle",
		Usage:       "A set of experimental verkle tree management commands",
		Description: "",
		Subcommands: []cli.Command{
			{
				Name:      "verify",
				Usage:     "verify the conversion of a MPT into a verkle tree",
				ArgsUsage: "<root>",
				Action:    verifyVerkle,
				Flags:     []cli.Flag{
					//utils.NetworkFlags,
					//utils.DatabasePathFlags,
				},
				Description: `
geth verkle verify <state-root>
This command takes a root commitment and attempts to rebuild the tree.
 `,
			},
			{
				Name:      "dump",
				Usage:     "Dump a verkle tree to a DOT file",
				ArgsUsage: "<root> <key1> [<key 2> ...]",
				Action:    expandVerkle,
				Flags:     []cli.Flag{
					//utils.NetworkFlags,
					//utils.DatabasePathFlags,
				},
				Description: `
geth verkle dump <state-root> <key 1> [<key 2> ...]
This command will produce a dot file representing the tree, rooted at <root>.
in which key1, key2, ... are expanded.
 `,
			},
		},
	}
)

// recurse into each child to ensure they can be loaded from the db. The tree isn't rebuilt
// (only its nodes are loaded) so there is no need to flush them, the garbage collector should
// take care of that for us.
func checkChildren(root verkle.VerkleNode, resolver verkle.NodeResolverFn) error {
	switch node := root.(type) {
	case *verkle.InternalNode:
		for i, child := range node.Children() {
			childC := child.Commit().Bytes()

			childS, err := resolver(childC[:])
			if bytes.Equal(childC[:], zero[:]) {
				continue
			}
			if err != nil {
				return fmt.Errorf("could not find child %x in db: %w", childC, err)
			}
			// depth is set to 0, the tree isn't rebuilt so it's not a problem
			childN, err := verkle.ParseNode(childS, 0, childC[:])
			if err != nil {
				return fmt.Errorf("decode error child %x in db: %w", child.Commitment().Bytes(), err)
			}
			if err := checkChildren(childN, resolver); err != nil {
				return fmt.Errorf("%x%w", i, err) // write the path to the erroring node
			}
		}
	case *verkle.LeafNode:
		// sanity check: ensure at least one value is non-zero

		for i := 0; i < verkle.NodeWidth; i++ {
			if len(node.Value(i)) != 0 {
				return nil
			}
		}
		return fmt.Errorf("Both balance and nonce are 0")
	case verkle.Empty:
		// nothing to do
	default:
		return fmt.Errorf("unsupported type encountered %v", root)
	}

	return nil
}

func verifyVerkle(ctx *cli.Context) error {
	stack, _ := makeConfigNode(ctx)
	defer stack.Close()

	chaindb := utils.MakeChainDatabase(ctx, stack, true)
	headBlock := rawdb.ReadHeadBlock(chaindb)
	if headBlock == nil {
		log.Error("Failed to load head block")
		return errors.New("no head block")
	}
	if ctx.NArg() > 1 {
		log.Error("Too many arguments given")
		return errors.New("too many arguments")
	}
	var (
		rootC common.Hash
		err   error
	)
	if ctx.NArg() == 1 {
		rootC, err = parseRoot(ctx.Args().First())
		if err != nil {
			log.Error("Failed to resolve state root", "error", err)
			return err
		}
		log.Info("Rebuilding the tree", "root", rootC)
	} else {
		rootC = headBlock.Root()
		log.Info("Rebuilding the tree", "root", rootC, "number", headBlock.NumberU64())
	}

	serializedRoot, err := chaindb.Get(rootC[:])
	if err != nil {
		return err
	}
	root, err := verkle.ParseNode(serializedRoot, 0, rootC[:])
	if err != nil {
		return err
	}

	if err := checkChildren(root, chaindb.Get); err != nil {
		log.Error("Could not rebuild the tree from the database", "err", err)
		return err
	}

	log.Info("Tree was rebuilt from the database")
	return nil
}

func expandVerkle(ctx *cli.Context) error {
	stack, _ := makeConfigNode(ctx)
	defer stack.Close()

	chaindb := utils.MakeChainDatabase(ctx, stack, true)
	var (
		rootC   common.Hash
		keylist [][]byte
		err     error
	)
	if ctx.NArg() >= 2 {
		rootC, err = parseRoot(ctx.Args().First())
		if err != nil {
			log.Error("Failed to resolve state root", "error", err)
			return err
		}
		keylist = make([][]byte, 0, len(ctx.Args())-1)
		args := ctx.Args()
		for i := range args[1:] {
			key, err := hex.DecodeString(args[i+1])
			log.Info("decoded key", "arg", args[i+1], "key", key)
			if err != nil {
				return fmt.Errorf("error decoding key #%d: %w", i+1, err)
			}
			keylist = append(keylist, key)
		}
		log.Info("Rebuilding the tree", "root", rootC)
	} else {
		return fmt.Errorf("usage: %s root key1 [key 2...]", ctx.App.Name)
	}

	serializedRoot, err := chaindb.Get(rootC[:])
	if err != nil {
		return err
	}
	root, err := verkle.ParseNode(serializedRoot, 0, rootC[:])
	if err != nil {
		return err
	}

	for i, key := range keylist {
		log.Info("Reading key", "index", i, "key", keylist[0])
		root.Get(key, chaindb.Get)
	}

	if err := os.WriteFile("dump.dot", []byte(verkle.ToDot(root)), 0600); err != nil {
		log.Error("Failed to dump file", "err", err)
	} else {
		log.Info("Tree was dumped to file", "file", "dump.dot")
	}
	return nil
}

func parseRoot(input string) (common.Hash, error) {
	var h common.Hash
	if err := h.UnmarshalText([]byte(input)); err != nil {
		return h, err
	}
	return h, nil
}
