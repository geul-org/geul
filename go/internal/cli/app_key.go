//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what app.key — a document address to its Key: a local Issuer key name first, then the archive's IDs and names
package cli

import (
	"strings"

	"github.com/geul-org/geul/go/pkg/core"
	"github.com/geul-org/geul/go/pkg/resolve"
)

// key resolves `<Issuer>/<문서>[@<버전>]`. An Issuer part that names a local
// key (localKeys) becomes that 16-word Issuer's ID; the rest is the data
// folder's Lookup (IDs only without a data folder).
func (a *app) key(s string) (resolve.Key, error) {
	is, rest, _ := strings.Cut(s, "/")
	keys, err := a.localKeys()
	if err != nil {
		return resolve.Key{}, err
	}
	for _, k := range keys {
		if k.Name == is {
			s = core.FormatID(k.Issuer) + "/" + rest
		}
	}
	d, _ := a.dir()
	return d.Lookup(s)
}
