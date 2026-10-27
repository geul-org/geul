//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what app.checkArchive — schema findings of the whole data folder, then every live version read back against its hash
package cli

import (
	"path/filepath"
)

// checkArchive checks every `.geul` under `.well-known/geul/` (checkPath) and
// then each document's versions (checkVersions). The walk has read the
// folder, so listing its Issuers and their documents does not fail.
func (a *app) checkArchive() ([]finding, error) {
	d, err := a.dir()
	if err != nil {
		return nil, err
	}
	out, err := checkPath(filepath.Join(d.Root, ".well-known", "geul"))
	if err != nil {
		return nil, err
	}
	issuers, _ := d.Issuers()
	for _, is := range issuers {
		got, _ := checkIssuer(d, is.ID)
		out = append(out, got...)
	}
	return out, nil
}
