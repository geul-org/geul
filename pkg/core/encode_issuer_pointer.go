//ff:func feature=core type=encoder control=sequence
//ff:what EncodeIssuerPointer — the ISSUER-POINTER Reference of an Issuer (Quantifier Identified)
package core

import "fmt"

// EncodeIssuerPointer builds the ISSUER-POINTER Reference (GEUL t002
// docs/reference/ISSUER-POINTER) for an Issuer of 1, 2, 4 or 16 words.
func EncodeIssuerPointer(issuer []Word) (Record, error) {
	if !IssuerCanonical(issuer) {
		return Record{}, fmt.Errorf("geul/core: ISSUER-POINTER needs a canonical 1/2/4/16-word Issuer (got %d words)", len(issuer))
	}
	return Record{Kind: KindReference, Codec: RefIssuerPointer << 3, Payload: append([]Word{}, issuer...)}, nil
}
