//ff:func feature=page type=helper control=selection
//ff:what refFields — the payload fields of a Reference by its codec: IMPORT ID, Issuer, document ID, version ID, target ID
package page

import "github.com/geul-org/gogeul/pkg/core"

// refFields splits a Reference payload of n words by codec c (the 4-bit CCCC):
// EXTERNAL-POINTER-* is IMPORT ID and target ID, EXTERNAL-ISSUER-* is Issuer,
// document ID and target ID, ISSUER-POINTER is the Issuer, DOC-POINTER is
// Issuer, document ID and version ID, and the declaration codecs are the
// target ID only.
// A field cut short by a short payload keeps what is there.
func refFields(c uint8, n int) []refField {
	switch {
	case c >= core.RefExternalPointer1Word && c <= core.RefExternalPointer8Word:
		w := min(core.WidthWords(c), n)
		return []refField{{"IMPORT ID", w}, {"대상 ID", n - w}}
	case c >= core.RefExternalIssuer1Word:
		w := min(core.ExternalIssuerWidth(c), n)
		d := min(4, n-w)
		return []refField{{"Issuer", w}, {"문서 ID", d}, {"대상 ID", n - w - d}}
	case c == core.RefIssuerPointer:
		return []refField{{"Issuer", n}}
	case c == core.RefDocPointer:
		w := max(n-6, 0)
		d := min(4, n-w)
		return []refField{{"Issuer", w}, {"문서 ID", d}, {"버전 ID", n - w - d}}
	default:
		return []refField{{"대상 ID", n}}
	}
}
