//ff:func feature=notation type=decoder control=selection
//ff:what formatReference — a Reference as a token: #n, Qn, Pn, a document address, an Issuer document's LOCAL, ISSUER-POINTER(…)
package notation

import "github.com/geul-org/geul/go/pkg/core"

// formatReference writes LOCAL, the Wikidata IDs, DOC-POINTER, ISSUER-POINTER
// and EXTERNAL-ISSUER-* with Quantifier Identified. Others are false.
func formatReference(r core.Record) (string, bool) {
	if r.Codec&7 != 0 {
		return "", false
	}
	switch c := r.Codec >> 3; {
	case c == core.RefLocal || c == core.RefWikidataQID || c == core.RefWikidataPID:
		return formatNumbered(c, r.Payload)
	case c == core.RefDocPointer:
		issuer, doc, ver, err := core.DecodeDocPointer(r)
		return docAddress(issuer, doc, uint32(ver[0])<<16|uint32(ver[1])), err == nil
	case c == core.RefIssuerPointer:
		issuer, err := core.DecodeIssuerPointer(r)
		return "ISSUER-POINTER(" + core.FormatID(issuer) + ")", err == nil
	case c >= core.RefExternalIssuer1Word:
		return formatExternalIssuer(r)
	default:
		return "", false
	}
}
