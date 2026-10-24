//ff:func feature=resolve type=decoder control=sequence
//ff:what issuerOf — the Issuer a Reference stands for: ISSUER-POINTER(Issuer)
package resolve

import "github.com/geul-org/geul/go/pkg/core"

// issuerOf returns the Issuer, in canonical hex, of an ISSUER-POINTER
// Reference, which stands for the Issuer itself (reference/ISSUER-POINTER).
func issuerOf(subj core.Record) (string, bool) {
	issuer, err := core.DecodeIssuerPointer(subj)
	if err != nil {
		return "", false
	}
	return core.FormatID(issuer), true
}
