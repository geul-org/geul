//ff:func feature=resolve type=decoder control=sequence
//ff:what docOf — the document a Reference stands for: EXTERNAL-ISSUER-*(Issuer, document ≠ 0, target 0)
package resolve

import "github.com/geul-org/geul/go/pkg/core"

// docOf returns the document, as a DocKey in canonical hex, of an
// EXTERNAL-ISSUER-* Reference whose target ID is 0 (the document itself,
// reference/LOCAL) and whose document ID is not 0 (that is the Issuer).
func docOf(subj core.Record) (DocKey, bool) {
	c := subj.Codec >> 3
	if subj.Kind != core.KindReference || c < core.RefExternalIssuer1Word || c > core.RefExternalIssuer16Word {
		return DocKey{}, false
	}
	w := core.ImportIssuerWidth(core.MetaImportDoc1Word + c&3)
	p := subj.Payload
	if len(p) != w+5 || p[w+4] != 0 || p[w]|p[w+1]|p[w+2]|p[w+3] == 0 {
		return DocKey{}, false
	}
	return DocKey{Issuer: core.FormatID(p[:w]), Doc: core.FormatID(p[w : w+4])}, true
}
