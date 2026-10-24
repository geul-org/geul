//ff:func feature=core type=decoder control=sequence
//ff:what DecodeIssuerPointer — the Issuer of an ISSUER-POINTER Reference (Quantifier Identified)
package core

import "fmt"

// DecodeIssuerPointer returns the Issuer of r, which must be an ISSUER-POINTER
// Reference with Quantifier Identified and a 1/2/4/16-word Issuer.
func DecodeIssuerPointer(r Record) ([]Word, error) {
	if r.Kind != KindReference || r.Codec != RefIssuerPointer<<3 || !issuerWidthOK(len(r.Payload)) {
		return nil, fmt.Errorf("geul/core: expected ISSUER-POINTER(Issuer 1/2/4/16 WORD), got kind %d codec 0x%02X, %d words", r.Kind, r.Codec, len(r.Payload))
	}
	return r.Payload, nil
}
