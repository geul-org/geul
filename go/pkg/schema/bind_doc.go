//ff:func feature=schema type=decoder control=sequence
//ff:what bindDoc — the document a BIND Subject names: EXTERNAL-ISSUER-*(Issuer, document, target 0)
package schema

import "github.com/geul-org/geul/go/pkg/core"

// bindDoc reads a version list BIND Subject, which must be an
// EXTERNAL-ISSUER-* Reference with target ID 0 (the document itself), and
// returns its document ID in canonical hex.
func bindDoc(subj core.Record) (string, bool) {
	c := subj.Codec >> 3
	if subj.Kind != core.KindReference || c < core.RefExternalIssuer1Word || c > core.RefExternalIssuer16Word {
		return "", false
	}
	w := core.ImportIssuerWidth(core.MetaImportDoc1Word + c&3)
	p := subj.Payload
	if len(p) != w+5 || p[w+4] != 0 {
		return "", false
	}
	return core.FormatID(p[w : w+4]), true
}
