// This work is subject to the CC0 1.0 Universal (CC0 1.0) Public Domain Dedication
// license. Its contents can be found at:
// http://creativecommons.org/publicdomain/zero/1.0/

package bindata

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"unicode"
)

// Translate reads assets from an input directory, converts them to Go code and
// writes new files to the output specified in the given configuration.
func Translate(c *Config) error {
	var toc []Asset

	// Ensure our configuration has sane values.
	if err := c.validate(); err != nil {
		return err
	}

	knownFuncs := make(map[string]int)
	visitedPaths := make(map[string]bool)
	// Locate all the assets.
	for _, input := range c.Input {
		if err := findFiles(input.Path, c.Prefix, input.Recursive, &toc, c.Ignore, knownFuncs, visitedPaths); err != nil {
			return err
		}
	}

	// Create output file.
	buf := new(bytes.Buffer)
	// Write the header. This makes e.g. Github ignore diffs in generated files.
	if _, err := fmt.Fprint(buf, "// Code generated by go-bindata. DO NOT EDIT.\n"); err != nil {
		return err
	}
	if _, err := fmt.Fprint(buf, "// sources:\n"); err != nil {
		return err
	}

	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	for _, asset := range toc {
		relative, err := filepath.Rel(wd, asset.Path)
		if err != nil {
			return err
		}
		relative = strings.TrimPrefix(relative, c.Prefix)
		relative = strings.TrimPrefix(relative, "/")
		if _, err = fmt.Fprintf(buf, "// %s (%s)\n", filepath.ToSlash(relative), bits(asset.Size)*byte_); err != nil {
			return err
		}
	}
	if _, err := fmt.Fprint(buf, "\n"); err != nil {
		return err
	}

	// Write build tags, if applicable.
	if len(c.Tags) > 0 {
		if _, err := fmt.Fprintf(buf, "// +build %s\n\n", c.Tags); err != nil {
			return err
		}
	}

	// Write package declaration.
	_, err = fmt.Fprintf(buf, "package %s\n\n", c.Package)
	if err != nil {
		return err
	}

	// Write assets.
	if c.Debug || c.Dev {
		if os.Getenv("GO_BINDATA_TEST") == "true" {
			// If we don't do this, people running the tests on different
			// machines get different git diffs.
			for i := range toc {
				toc[i].Path = strings.Replace(toc[i].Path, wd, "/test", 1)
			}
		}
		err = writeDebugFunctions(buf, c, toc)
	} else {
		err = writeReleaseFunctions(buf, c, toc)
	}
	if err != nil {
		return err
	}

	// Write table of contents
	if err := writeTOC(buf, toc); err != nil {
		return err
	}
	_, err = fmt.Fprintf(buf, `// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = %t

`, c.Debug)
	if err != nil {
		return err
	}
	// Write hierarchical tree of assets
	if err := writeTOCTree(buf, toc); err != nil {
		return err
	}

	// Write restore procedure
	if err := writeRestore(buf); err != nil {
		return err
	}

	return diffAndWrite(c.Output, buf.Bytes(), 0666)
}

// findFiles recursively finds all the file paths in the given directory tree.
// They are added to the given map as keys. Values will be safe function names
// for each file, which will be used when generating the output code.
func findFiles(dirOrFile, prefix string, recursive bool, toc *[]Asset, ignore []*regexp.Regexp, knownFuncs map[string]int, visitedPaths map[string]bool) error {
	dirOrFile = filepath.Clean(dirOrFile)
	// confusingly, if a prefix exists this is the absolute path; if a prefix
	// doesn't exist it may be an absolute path, or it may be a relative path.
	var dirOrFilePath string
	if len(prefix) > 0 {
		var err error
		dirOrFilePath, err = filepath.Abs(dirOrFile)
		if err != nil {
			return err
		}
		prefix, err = filepath.Abs(prefix)
		if err != nil {
			return err
		}
		prefix = filepath.ToSlash(prefix)
	} else {
		dirOrFilePath = dirOrFile
	}

	dirOrFileFI, err := os.Stat(dirOrFilePath)
	if err != nil {
		return err
	}

	var list []os.FileInfo
	// if prefix is non-empty this is the absolute path, otherwise it may be
	// a relative path.
	var dirpath string
	if dirOrFileFI.IsDir() {
		dirpath = dirOrFilePath
		visitedPaths[dirpath] = true
		fd, err := os.Open(dirpath)
		if err != nil {
			return err
		}

		list, err = fd.Readdir(0)
		if err != nil {
			return err
		}

		// Sort to make output stable between invocations
		sort.Slice(list, func(i, j int) bool {
			return list[i].Name() < list[j].Name()
		})
		fd.Close()
	} else {
		dirpath = filepath.Dir(dirOrFilePath)
		list = []os.FileInfo{dirOrFileFI}
	}

	for _, entry := range list {
		var asset Asset
		asset.Path = filepath.Join(dirpath, entry.Name())
		pathWithSlashes := filepath.ToSlash(asset.Path)

		shouldIgnore := false
		for _, re := range ignore {
			if re.MatchString(asset.Path) {
				shouldIgnore = true
				break
			}
		}
		if shouldIgnore {
			continue
		}

		if entry.IsDir() {
			if recursive {
				recursivePath := filepath.Join(dirOrFile, entry.Name())
				visitedPaths[asset.Path] = true
				findFiles(recursivePath, prefix, recursive, toc, ignore, knownFuncs, visitedPaths)
			}
			continue
		} else if entry.Mode()&os.ModeSymlink == os.ModeSymlink {
			var linkPath string
			if linkPath, err = os.Readlink(asset.Path); err != nil {
				return err
			}
			if !filepath.IsAbs(linkPath) {
				if linkPath, err = filepath.Abs(dirpath + "/" + linkPath); err != nil {
					return err
				}
			}
			if _, ok := visitedPaths[linkPath]; !ok {
				visitedPaths[linkPath] = true
				findFiles(asset.Path, prefix, recursive, toc, ignore, knownFuncs, visitedPaths)
			}
			continue
		}

		if strings.HasPrefix(pathWithSlashes, prefix) {
			asset.Name = strings.TrimPrefix(pathWithSlashes, prefix)
		} else {
			// File or directory isn't inside of the prefix list
			if dirOrFileFI.IsDir() {
				asset.Name = filepath.Join(dirOrFile, entry.Name())
			} else {
				asset.Name = dirOrFile
			}
		}

		// If we have a leading slash, get rid of it.
		asset.Name = strings.TrimPrefix(asset.Name, "/")

		// This shouldn't happen.
		if len(asset.Name) == 0 {
			return fmt.Errorf("invalid file: %v", asset.Path)
		}

		asset.Func = safeFunctionName(asset.Name, knownFuncs)
		asset.Path, err = filepath.Abs(asset.Path)
		if err != nil {
			return err
		}
		asset.Size = entry.Size()
		*toc = append(*toc, asset)
	}

	return nil
}

var regFuncName = regexp.MustCompile(`[^a-zA-Z0-9_]`)

// safeFunctionName converts the given name into a name which qualifies as a
// valid function identifier. It also compares against a known list of functions
// (if one is provided) to prevent conflict based on name translation.
func safeFunctionName(name string, knownFuncs map[string]int) string {
	var inBytes, outBytes []byte
	var toUpper bool

	name = strings.ToLower(name)
	inBytes = []byte(name)

	for i := 0; i < len(inBytes); i++ {
		if regFuncName.Match([]byte{inBytes[i]}) {
			toUpper = true
		} else if toUpper {
			outBytes = append(outBytes, []byte(strings.ToUpper(string(inBytes[i])))...)
			toUpper = false
		} else {
			outBytes = append(outBytes, inBytes[i])
		}
	}

	name = string(outBytes)

	// Identifier can't start with a digit.
	if unicode.IsDigit(rune(name[0])) {
		name = "_" + name
	}

	if num, ok := knownFuncs[name]; ok {
		knownFuncs[name] = num + 1
		name = fmt.Sprintf("%s%d", name, num)
	} else if knownFuncs != nil {
		knownFuncs[name] = 2
	}

	return name
}
