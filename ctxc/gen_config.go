// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package ctxc

import (
	"time"

	"github.com/CortexFoundation/CortexTheseus/common"
	"github.com/CortexFoundation/CortexTheseus/consensus/cuckoo"
	"github.com/CortexFoundation/CortexTheseus/core"
	"github.com/CortexFoundation/CortexTheseus/core/txpool"
	"github.com/CortexFoundation/CortexTheseus/ctxc/downloader"
	"github.com/CortexFoundation/CortexTheseus/ctxc/gasprice"
	"github.com/CortexFoundation/CortexTheseus/miner"
	"github.com/CortexFoundation/CortexTheseus/params"
)

// MarshalTOML marshals as TOML.
func (c Config) MarshalTOML() (interface{}, error) {
	type Config struct {
		Genesis                 *core.Genesis `toml:",omitempty"`
		NetworkId               uint64
		SyncMode                downloader.SyncMode
		DiscoveryURLs           []string
		NoPruning               bool
		NoPrefetch              bool
		TxLookupLimit           uint64                 `toml:",omitempty"`
		Whitelist               map[uint64]common.Hash `toml:"-"`
		SkipBcVersionCheck      bool                   `toml:"-"`
		DatabaseHandles         int                    `toml:"-"`
		DatabaseCache           int
		DatabaseFreezer         string
		TrieCleanCache          int
		TrieCleanCacheJournal   string        `toml:",omitempty"`
		TrieCleanCacheRejournal time.Duration `toml:",omitempty"`
		TrieDirtyCache          int
		TrieTimeout             time.Duration
		SnapshotCache           int
		Preimages               bool
		Miner                   miner.Config
		Coinbase                common.Address `toml:",omitempty"`
		InferDeviceType         string
		InferDeviceId           int
		SynapseTimeout          int
		InferMemoryUsage        int64
		Cuckoo                  cuckoo.Config
		TxPool                  txpool.Config
		GPO                     gasprice.Config
		EnablePreimageRecording bool
		InferURI                string
		StorageDir              string
		DocRoot                 string                         `toml:"-"`
		RPCGasCap               uint64                         `toml:",omitempty"`
		RPCTxFeeCap             float64                        `toml:",omitempty"`
		Checkpoint              *params.TrustedCheckpoint      `toml:",omitempty"`
		CheckpointOracle        *params.CheckpointOracleConfig `toml:",omitempty"`
		Viper                   bool
	}
	var enc Config
	enc.Genesis = c.Genesis
	enc.NetworkId = c.NetworkId
	enc.SyncMode = c.SyncMode
	enc.DiscoveryURLs = c.DiscoveryURLs
	enc.NoPruning = c.NoPruning
	enc.NoPrefetch = c.NoPrefetch
	enc.TxLookupLimit = c.TxLookupLimit
	enc.Whitelist = c.Whitelist
	enc.SkipBcVersionCheck = c.SkipBcVersionCheck
	enc.DatabaseHandles = c.DatabaseHandles
	enc.DatabaseCache = c.DatabaseCache
	enc.DatabaseFreezer = c.DatabaseFreezer
	enc.TrieCleanCache = c.TrieCleanCache
	enc.TrieCleanCacheJournal = c.TrieCleanCacheJournal
	enc.TrieCleanCacheRejournal = c.TrieCleanCacheRejournal
	enc.TrieDirtyCache = c.TrieDirtyCache
	enc.TrieTimeout = c.TrieTimeout
	enc.SnapshotCache = c.SnapshotCache
	enc.Preimages = c.Preimages
	enc.Miner = c.Miner
	enc.Coinbase = c.Coinbase
	enc.InferDeviceType = c.InferDeviceType
	enc.InferDeviceId = c.InferDeviceId
	enc.SynapseTimeout = c.SynapseTimeout
	enc.InferMemoryUsage = c.InferMemoryUsage
	enc.Cuckoo = c.Cuckoo
	enc.TxPool = c.TxPool
	enc.GPO = c.GPO
	enc.EnablePreimageRecording = c.EnablePreimageRecording
	enc.InferURI = c.InferURI
	enc.StorageDir = c.StorageDir
	enc.DocRoot = c.DocRoot
	enc.RPCGasCap = c.RPCGasCap
	enc.RPCTxFeeCap = c.RPCTxFeeCap
	enc.Checkpoint = c.Checkpoint
	enc.CheckpointOracle = c.CheckpointOracle
	enc.Viper = c.Viper
	return &enc, nil
}

// UnmarshalTOML unmarshals from TOML.
func (c *Config) UnmarshalTOML(unmarshal func(interface{}) error) error {
	type Config struct {
		Genesis                 *core.Genesis `toml:",omitempty"`
		NetworkId               *uint64
		SyncMode                *downloader.SyncMode
		DiscoveryURLs           []string
		NoPruning               *bool
		NoPrefetch              *bool
		TxLookupLimit           *uint64                `toml:",omitempty"`
		Whitelist               map[uint64]common.Hash `toml:"-"`
		SkipBcVersionCheck      *bool                  `toml:"-"`
		DatabaseHandles         *int                   `toml:"-"`
		DatabaseCache           *int
		DatabaseFreezer         *string
		TrieCleanCache          *int
		TrieCleanCacheJournal   *string        `toml:",omitempty"`
		TrieCleanCacheRejournal *time.Duration `toml:",omitempty"`
		TrieDirtyCache          *int
		TrieTimeout             *time.Duration
		SnapshotCache           *int
		Preimages               *bool
		Miner                   *miner.Config
		Coinbase                *common.Address `toml:",omitempty"`
		InferDeviceType         *string
		InferDeviceId           *int
		SynapseTimeout          *int
		InferMemoryUsage        *int64
		Cuckoo                  *cuckoo.Config
		TxPool                  *txpool.Config
		GPO                     *gasprice.Config
		EnablePreimageRecording *bool
		InferURI                *string
		StorageDir              *string
		DocRoot                 *string                        `toml:"-"`
		RPCGasCap               *uint64                        `toml:",omitempty"`
		RPCTxFeeCap             *float64                       `toml:",omitempty"`
		Checkpoint              *params.TrustedCheckpoint      `toml:",omitempty"`
		CheckpointOracle        *params.CheckpointOracleConfig `toml:",omitempty"`
		Viper                   *bool
	}
	var dec Config
	if err := unmarshal(&dec); err != nil {
		return err
	}
	if dec.Genesis != nil {
		c.Genesis = dec.Genesis
	}
	if dec.NetworkId != nil {
		c.NetworkId = *dec.NetworkId
	}
	if dec.SyncMode != nil {
		c.SyncMode = *dec.SyncMode
	}
	if dec.DiscoveryURLs != nil {
		c.DiscoveryURLs = dec.DiscoveryURLs
	}
	if dec.NoPruning != nil {
		c.NoPruning = *dec.NoPruning
	}
	if dec.NoPrefetch != nil {
		c.NoPrefetch = *dec.NoPrefetch
	}
	if dec.TxLookupLimit != nil {
		c.TxLookupLimit = *dec.TxLookupLimit
	}
	if dec.Whitelist != nil {
		c.Whitelist = dec.Whitelist
	}
	if dec.SkipBcVersionCheck != nil {
		c.SkipBcVersionCheck = *dec.SkipBcVersionCheck
	}
	if dec.DatabaseHandles != nil {
		c.DatabaseHandles = *dec.DatabaseHandles
	}
	if dec.DatabaseCache != nil {
		c.DatabaseCache = *dec.DatabaseCache
	}
	if dec.DatabaseFreezer != nil {
		c.DatabaseFreezer = *dec.DatabaseFreezer
	}
	if dec.TrieCleanCache != nil {
		c.TrieCleanCache = *dec.TrieCleanCache
	}
	if dec.TrieCleanCacheJournal != nil {
		c.TrieCleanCacheJournal = *dec.TrieCleanCacheJournal
	}
	if dec.TrieCleanCacheRejournal != nil {
		c.TrieCleanCacheRejournal = *dec.TrieCleanCacheRejournal
	}
	if dec.TrieDirtyCache != nil {
		c.TrieDirtyCache = *dec.TrieDirtyCache
	}
	if dec.TrieTimeout != nil {
		c.TrieTimeout = *dec.TrieTimeout
	}
	if dec.SnapshotCache != nil {
		c.SnapshotCache = *dec.SnapshotCache
	}
	if dec.Preimages != nil {
		c.Preimages = *dec.Preimages
	}
	if dec.Miner != nil {
		c.Miner = *dec.Miner
	}
	if dec.Coinbase != nil {
		c.Coinbase = *dec.Coinbase
	}
	if dec.InferDeviceType != nil {
		c.InferDeviceType = *dec.InferDeviceType
	}
	if dec.InferDeviceId != nil {
		c.InferDeviceId = *dec.InferDeviceId
	}
	if dec.SynapseTimeout != nil {
		c.SynapseTimeout = *dec.SynapseTimeout
	}
	if dec.InferMemoryUsage != nil {
		c.InferMemoryUsage = *dec.InferMemoryUsage
	}
	if dec.Cuckoo != nil {
		c.Cuckoo = *dec.Cuckoo
	}
	if dec.TxPool != nil {
		c.TxPool = *dec.TxPool
	}
	if dec.GPO != nil {
		c.GPO = *dec.GPO
	}
	if dec.EnablePreimageRecording != nil {
		c.EnablePreimageRecording = *dec.EnablePreimageRecording
	}
	if dec.InferURI != nil {
		c.InferURI = *dec.InferURI
	}
	if dec.StorageDir != nil {
		c.StorageDir = *dec.StorageDir
	}
	if dec.DocRoot != nil {
		c.DocRoot = *dec.DocRoot
	}
	if dec.RPCGasCap != nil {
		c.RPCGasCap = *dec.RPCGasCap
	}
	if dec.RPCTxFeeCap != nil {
		c.RPCTxFeeCap = *dec.RPCTxFeeCap
	}
	if dec.Checkpoint != nil {
		c.Checkpoint = dec.Checkpoint
	}
	if dec.CheckpointOracle != nil {
		c.CheckpointOracle = dec.CheckpointOracle
	}
	if dec.Viper != nil {
		c.Viper = *dec.Viper
	}
	return nil
}
