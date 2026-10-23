//ff:func feature=schema type=decoder control=iteration dimension=1
//ff:what listIssuer — the Issuer of a version list document's DOC-ID in canonical hex
package schema

import "github.com/geul-org/geul/go/pkg/core"

// listIssuer returns the Issuer part of the HEAD's DOC-ID, "" when absent.
func listIssuer(head []rec) string {
	for _, r := range head {
		if r.R.Codec == core.MetaDocID && len(r.R.Payload) > 4 {
			return core.FormatID(r.R.Payload[:len(r.R.Payload)-4])
		}
	}
	return ""
}
