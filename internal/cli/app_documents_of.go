//ff:func feature=cli type=command control=sequence
//ff:what app.documentsOf — the documents of an Issuer given by ID, name or local key name
package cli

import "github.com/geul-org/gogeul/pkg/resolve"

// documentsOf resolves issuer (key, with a placeholder document) and lists
// its documents.
func (a *app) documentsOf(d resolve.Dir, issuer string) ([]resolve.Named, error) {
	k, err := a.key(issuer + "/0x0000000000000000")
	if err != nil {
		return nil, err
	}
	return d.Documents(k.Issuer)
}
