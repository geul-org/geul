//ff:func feature=core type=helper control=selection
//ff:what ExternalIssuerWidth — the Issuer word count of an EXTERNAL-ISSUER-* Reference codec, 0 otherwise
package core

// ExternalIssuerWidth returns the Issuer width that the 4-bit Reference codec
// c of EXTERNAL-ISSUER-1/2/4/16WORD declares (GEUL t002 OVERVIEW Reference):
// its low 2 bits, 00=1, 01=2, 10=4, 11=16 words. Any other codec returns 0.
func ExternalIssuerWidth(c uint8) int {
	switch c {
	case RefExternalIssuer1Word:
		return 1
	case RefExternalIssuer2Word:
		return 2
	case RefExternalIssuer4Word:
		return 4
	case RefExternalIssuer16Word:
		return 16
	default:
		return 0
	}
}
