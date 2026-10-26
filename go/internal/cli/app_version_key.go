//ff:func feature=cli type=helper control=sequence
//ff:what app.versionKey — a document address that must name a version and a data folder to act in
package cli

import (
	"errors"

	"github.com/geul-org/geul/go/pkg/resolve"
)

// versionKey resolves addr (key) and requires @VERSION and a data folder,
// for commands that change one version's state.
func (a *app) versionKey(addr string) (resolve.Key, error) {
	if _, err := a.dir(); err != nil {
		return resolve.Key{}, err
	}
	k, err := a.key(addr)
	if err == nil && k.Version == 0 {
		err = errors.New("give the version: DOCUMENT@VERSION")
	}
	return k, err
}
