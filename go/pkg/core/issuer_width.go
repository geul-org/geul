//ff:func feature=core type=helper control=selection
//ff:what IssuerWidth — the Issuer word count of an ISSUER codec, 0 for a direct codec
package core

// IssuerWidth returns how many Issuer words the Codec extension TAG of a k/codec
// record carries: 1/2/4/16 for the ISSUER codecs of Value and Edge, and 0 for
// every direct codec. Reference and Meta have no ISSUER codes (GEUL t002
// README), so they always return 0.
func IssuerWidth(k Kind, codec uint8) int {
	if k != KindValue && k != KindEdge {
		return 0
	}
	switch codec {
	case CodecIssuer1Word:
		return 1
	case CodecIssuer2Word:
		return 2
	case CodecIssuer4Word:
		return 4
	case CodecIssuer16Word:
		return 16
	default:
		return 0
	}
}
