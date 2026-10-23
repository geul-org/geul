//ff:func feature=core type=helper control=selection
//ff:what issuerRefCodec — the EXTERNAL-ISSUER-* Reference codec (Quantifier Identified) for an Issuer width; false for other widths
package core

// issuerRefCodec returns the record codec of EXTERNAL-ISSUER-1/2/4/16WORD with
// Quantifier Identified for an Issuer of w words.
func issuerRefCodec(w int) (uint8, bool) {
	switch w {
	case 1:
		return RefExternalIssuer1Word << 3, true
	case 2:
		return RefExternalIssuer2Word << 3, true
	case 4:
		return RefExternalIssuer4Word << 3, true
	case 16:
		return RefExternalIssuer16Word << 3, true
	default:
		return 0, false
	}
}
