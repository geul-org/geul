//ff:func feature=resolve type=helper control=sequence
//ff:what Key.DocKey — the version-free identity of a Key
package resolve

import "github.com/geul-org/gogeul/pkg/core"

// DocKey returns k's Issuer and document ID in canonical hex.
func (k Key) DocKey() DocKey {
	return DocKey{Issuer: core.FormatID(k.Issuer), Doc: core.FormatID(k.Doc)}
}
