//ff:func feature=core type=decoder control=sequence
//ff:what DecodeDocPointer — the Issuer, document ID and version ID of a DOC-POINTER Reference (Quantifier Identified)
package core

import "fmt"

// DecodeDocPointer splits r, which must be a DOC-POINTER Reference with
// Quantifier Identified, into its Issuer (the payload less the 6 fixed words,
// 1/2/4/16 words), its 4-word document ID (not 0) and its 2-word version ID.
func DecodeDocPointer(r Record) (issuer, doc, version []Word, err error) {
	w := len(r.Payload) - 6
	if r.Kind != KindReference || r.Codec != RefDocPointer<<3 || w < 1 || !IssuerCanonical(r.Payload[:w]) {
		return nil, nil, nil, fmt.Errorf("geul/core: expected DOC-POINTER(Issuer, document ID, version ID), got kind %d codec 0x%02X, %d words", r.Kind, r.Codec, len(r.Payload))
	}
	p := r.Payload
	if p[w]|p[w+1]|p[w+2]|p[w+3] == 0 {
		return nil, nil, nil, fmt.Errorf("geul/core: DOC-POINTER document ID 0")
	}
	return p[:w], p[w : w+4], p[w+4:], nil
}
