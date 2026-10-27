//ff:func feature=core type=decoder control=sequence
//ff:what DecodeDocID — the Issuer and document ID of Meta DOC-ID
package core

import "fmt"

// DecodeDocID reads Meta DOC-ID: exactly one DOC-POINTER Reference with
// Quantifier Identified and version ID 0.
func DecodeDocID(r Record) (issuer, doc []Word, err error) {
	if r.Kind != KindMeta || r.Codec != MetaDocID {
		return nil, nil, fmt.Errorf("geul/core: not a DOC-ID (kind %d codec 0x%02X)", r.Kind, r.Codec)
	}
	ref, n, err := DecodeRecord(r.Payload)
	if err != nil || n != len(r.Payload) {
		return nil, nil, fmt.Errorf("geul/core: DOC-ID payload must be one Reference record")
	}
	issuer, doc, ver, err := DecodeDocPointer(ref)
	if err != nil || ver[0]|ver[1] != 0 {
		return nil, nil, fmt.Errorf("geul/core: DOC-ID must be DOC-POINTER(Issuer, document ID, 0)")
	}
	return issuer, doc, nil
}
