//ff:func feature=schema type=decoder control=sequence
//ff:what bindDoc — the document ID a BIND Subject names: EXTERNAL-ISSUER of the list's Issuer with target 0
package schema

import "github.com/geul-org/geul/go/pkg/core"

// bindDoc reads a version list BIND Subject, which must be an
// EXTERNAL-ISSUER-* Reference of the list's own Issuer whose target ID is 0
// (the document itself), and returns the document ID in canonical hex.
func bindDoc(subj core.Record, issuer string) (string, bool) {
	c := subj.Codec >> 3
	if subj.Kind != core.KindReference || c < core.RefExternalIssuer1Word || c > core.RefExternalIssuer16Word {
		return "", false
	}
	w := core.ImportIssuerWidth(core.MetaImportDoc1Word + c&3)
	p := subj.Payload
	if len(p) != w+5 || p[w+4] != 0 || core.FormatID(p[:w]) != issuer {
		return "", false
	}
	return core.FormatID(p[w : w+4]), true
}
