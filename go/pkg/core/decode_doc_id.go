//ff:func feature=core type=decoder control=sequence
//ff:what DecodeDocID — the Issuer and document ID of Meta DOC-ID
package core

import "fmt"

// DecodeDocID reads Meta DOC-ID: exactly one EXTERNAL-ISSUER-* Reference with
// Quantifier Identified whose target ID is 0 and whose document ID is not 0.
func DecodeDocID(r Record) (issuer, doc []Word, err error) {
	if r.Kind != KindMeta || r.Codec != MetaDocID {
		return nil, nil, fmt.Errorf("geul/core: not a DOC-ID (kind %d codec 0x%02X)", r.Kind, r.Codec)
	}
	ref, n, err := DecodeRecord(r.Payload)
	if err != nil || n != len(r.Payload) {
		return nil, nil, fmt.Errorf("geul/core: DOC-ID payload must be one Reference record")
	}
	w := ImportIssuerWidth(MetaImportDoc1Word + (ref.Codec>>3)&3)
	c, _ := issuerRefCodec(w)
	p := ref.Payload
	if ref.Kind != KindReference || ref.Codec != c || len(p) != w+5 || p[w+4] != 0 || p[w]|p[w+1]|p[w+2]|p[w+3] == 0 {
		return nil, nil, fmt.Errorf("geul/core: DOC-ID must be EXTERNAL-ISSUER-*(Issuer, document ID, 0)")
	}
	return p[:w], p[w : w+4], nil
}
