//ff:func feature=core type=decoder control=sequence
//ff:what DecodeRecord — decode one GEUL record from the head of ws, returning it and the words consumed
package core

import (
	"errors"
	"fmt"
)

// DecodeRecord decodes the record at the head of ws (GEUL t002 README). It
// enforces the 000 prefix and the canonical Length, slices the Codec extension
// TAG off an ISSUER codec, and returns the Record with the total word count, so
// a caller can skip records without a codebook.
func DecodeRecord(ws []Word) (Record, int, error) {
	if len(ws) < 1 {
		return Record{}, 0, errors.New("geul/core: record requires a TAG WORD")
	}
	tag := ws[0]
	if tag>>13 != 0 {
		return Record{}, 0, fmt.Errorf("geul/core: TAG WORD 0x%04X prefix is not 000", uint16(tag))
	}
	r := Record{Kind: Kind((tag >> 11) & 3), Codec: uint8(tag & 0x7F)}
	total, head, err := decodeLength(ws, uint8((tag>>7)&0xF))
	if err != nil {
		return Record{}, 0, err
	}
	if total < head || len(ws) < total {
		return Record{}, 0, errors.New("geul/core: record Length runs past the input")
	}
	body := ws[head:total]
	w := IssuerWidth(r.Kind, r.Codec)
	if w > 0 && len(body) < w+1 {
		return Record{}, 0, errors.New("geul/core: Codec extension TAG truncated")
	}
	if w > 0 {
		r.Issuer = append([]Word(nil), body[:w]...)
		r.Format = uint16(body[w])
		body = body[w+1:]
	}
	r.Payload = body
	return r, total, nil
}
