//ff:func feature=cli type=helper control=sequence
//ff:what app.dir — the data folder as a local archive; an error when none is set or found
package cli

import (
	"errors"

	"github.com/geul-org/geul/go/pkg/resolve"
)

// dir returns the archive at GEULDATA.
func (a *app) dir() (resolve.Dir, error) {
	root := a.cfg.Get("GEULDATA")
	if root == "" {
		return resolve.Dir{}, errors.New("no data folder: give --data, set GEULDATA, or run inside a folder holding .well-known/geul/")
	}
	return resolve.Dir{Root: root}, nil
}
