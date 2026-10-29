//ff:func feature=cli type=helper control=sequence
//ff:what app.dir — the data folder as a local archive (pack sets or directory tree); an error when none is set or found
package cli

import (
	"errors"

	"github.com/geul-org/gogeul/pkg/resolve"
)

// dir returns the archive at GEULDATA, backed by its pack archive when it has one.
func (a *app) dir() (resolve.Dir, error) {
	root := a.cfg.Get("GEULDATA")
	if root == "" {
		return resolve.Dir{}, errors.New("no data folder: give --data, set GEULDATA, or run inside a folder holding .well-known/geul/")
	}
	arch, err := a.archive(root)
	if err != nil {
		return resolve.Dir{}, err
	}
	d := resolve.Dir{Root: root}
	if arch != nil {
		d.Store = arch
	}
	return d, nil
}
