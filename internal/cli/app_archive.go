//ff:func feature=cli type=helper control=sequence
//ff:what app.archive — the data folder's pack archive (opened once per run), nil for a directory archive
package cli

import (
	"context"
	"path/filepath"

	"github.com/geul-org/gogeul/pkg/pack"
)

// archive opens root's archive the first time and keeps it, so staged writes
// stay visible to every later read of the run. Its writes use GEULCOMPRESS.
func (a *app) archive(root string) (*pack.Archive, error) {
	if a.arch != nil || a.dirs {
		return a.arch, nil
	}
	wk := filepath.Join(root, ".well-known", "geul")
	if !usePack(wk, a.cfg.Get("GEULSTORE")) {
		a.dirs = true
		return nil, nil
	}
	o, err := a.writeOptions("")
	if err != nil {
		return nil, err
	}
	arch, err := pack.OpenDir(context.Background(), wk, o)
	a.arch = arch
	return arch, err
}
