//ff:func feature=core type=encoder control=sequence
//ff:what EncodeDocPointer — the DOC-POINTER Reference of one document version (Quantifier Identified)
package core

import (
	"fmt"
	"slices"
)

// EncodeDocPointer builds the DOC-POINTER Reference (GEUL t002
// docs/reference/DOC-POINTER): a canonical 1/2/4/16-word Issuer (IssuerCanonical), a 4-word document ID
// that is not 0 and a 2-word version ID, where version 0 is the latest.
func EncodeDocPointer(issuer, doc, version []Word) (Record, error) {
	if !IssuerCanonical(issuer) || len(doc) != 4 || len(version) != 2 {
		return Record{}, fmt.Errorf("geul/core: DOC-POINTER needs a canonical 1/2/4/16-word Issuer, a 4-word document ID and a 2-word version ID (got %d, %d, %d)", len(issuer), len(doc), len(version))
	}
	if doc[0]|doc[1]|doc[2]|doc[3] == 0 {
		return Record{}, fmt.Errorf("geul/core: DOC-POINTER document ID 0 is the Issuer itself (ISSUER-POINTER)")
	}
	return Record{Kind: KindReference, Codec: RefDocPointer << 3, Payload: slices.Concat(issuer, doc, version)}, nil
}
