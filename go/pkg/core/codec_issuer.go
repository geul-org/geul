//ff:type feature=core type=model
//ff:what ISSUER codec constants — the four Value/Edge codecs whose Codec extension TAG carries an Issuer
package core

// The ISSUER codecs of the Value and Edge kinds (GEUL t002 README). Their Codec
// extension TAG holds the Issuer (1/2/4/16 words by codec) and a 1-word format
// number defined by that Issuer.
const (
	// CodecIssuer1Word carries a 1-word Issuer.
	CodecIssuer1Word uint8 = 0x7C
	// CodecIssuer2Word carries a 2-word Issuer.
	CodecIssuer2Word uint8 = 0x7D
	// CodecIssuer4Word carries a 4-word Issuer.
	CodecIssuer4Word uint8 = 0x7E
	// CodecIssuer16Word carries a 16-word Issuer.
	CodecIssuer16Word uint8 = 0x7F
)
