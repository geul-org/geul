//ff:func feature=cli type=helper control=sequence
//ff:what app.keysDir — where the private keys of self-issued Issuers live: outside the data folder, in the user's config folder
package cli

import "path/filepath"

// keysDir is `<config>/geul/keys`. Keys never go in the data folder, which is
// published (README Issuer 명령).
func (a *app) keysDir() string {
	return filepath.Join(a.env.ConfigDir, "geul", "keys")
}
