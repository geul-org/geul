//ff:func feature=cli type=helper control=sequence
//ff:what openArchive — the pack archive a folder holds (data folder or its .well-known/geul); an error for a directory archive
package cli

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/geul-org/gogeul/pkg/pack"
)

// openArchive opens the pack sets under dir; o applies to writes (compact).
func openArchive(dir string, o pack.BuildOptions) (*pack.Archive, error) {
	root, err := treeRoot(dir)
	if err != nil {
		return nil, err
	}
	wk := filepath.Join(root, ".well-known", "geul")
	if !pack.IsPacked(wk) {
		return nil, fmt.Errorf("%s: not a pack archive (convert it with geul pack build)", wk)
	}
	return pack.OpenDir(context.Background(), wk, o)
}
