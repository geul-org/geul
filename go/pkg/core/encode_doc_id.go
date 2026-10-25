//ff:func feature=core type=encoder control=sequence
//ff:what EncodeDocID — Meta DOC-ID: one DOC-POINTER(Issuer, document ID, version 0) Reference
package core

import "fmt"

// EncodeDocID builds Meta DOC-ID (GEUL t002 docs/meta/declaration/DOC-ID):
// its payload is the Reference DOC-POINTER(issuer, doc, 0), the document
// itself. issuer must be canonical and doc 4 words (0 is the Issuer's term
// definition index document). The
// version is 0 because the version list, not the document, assigns versions.
func EncodeDocID(issuer, doc []Word) (Record, error) {
	ptr, err := EncodeDocPointer(issuer, doc, []Word{0, 0})
	if err != nil {
		return Record{}, fmt.Errorf("geul/core: DOC-ID: %w", err)
	}
	ref, _ := EncodeRecord(ptr)
	return Record{Kind: KindMeta, Codec: MetaDocID, Payload: ref}, nil
}
