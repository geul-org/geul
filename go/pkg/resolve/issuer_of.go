//ff:func feature=resolve type=decoder control=sequence
//ff:what issuerOf — the Issuer a Reference stands for: EXTERNAL-ISSUER-*(Issuer, document 0, target 0)
package resolve

import "github.com/geul-org/geul/go/pkg/core"

// issuerOf returns the Issuer, in canonical hex, of an EXTERNAL-ISSUER-*
// Reference whose document ID and target ID are both 0: the Issuer's version
// list document itself, which stands for the Issuer (OVERVIEW Reference).
func issuerOf(subj core.Record) (string, bool) {
	c := subj.Codec >> 3
	if subj.Kind != core.KindReference || c < core.RefExternalIssuer1Word || c > core.RefExternalIssuer16Word {
		return "", false
	}
	w := core.ImportIssuerWidth(core.MetaImportDoc1Word + c&3)
	p := subj.Payload
	if len(p) != w+5 || p[w]|p[w+1]|p[w+2]|p[w+3]|p[w+4] != 0 {
		return "", false
	}
	return core.FormatID(p[:w]), true
}
