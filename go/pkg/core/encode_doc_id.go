//ff:func feature=core type=encoder control=sequence
//ff:what EncodeDocID — Meta DOC-ID: one EXTERNAL-ISSUER-*(Issuer, document ID, 0) Reference
package core

import "fmt"

// EncodeDocID builds Meta DOC-ID (GEUL t002 docs/meta/declaration/DOC-ID):
// its payload is the Reference EXTERNAL-ISSUER-*(issuer, doc, 0), the document
// itself. issuer must be 1, 2, 4 or 16 words and doc 4 words.
func EncodeDocID(issuer, doc []Word) (Record, error) {
	c, ok := issuerRefCodec(len(issuer))
	if !ok || len(doc) != 4 {
		return Record{}, fmt.Errorf("geul/core: DOC-ID needs a 1/2/4/16-word Issuer and a 4-word document ID (got %d, %d)", len(issuer), len(doc))
	}
	ref, _ := EncodeRecord(Record{Kind: KindReference, Codec: c, Payload: append(append(append([]Word{}, issuer...), doc...), 0)})
	return Record{Kind: KindMeta, Codec: MetaDocID, Payload: ref}, nil
}
