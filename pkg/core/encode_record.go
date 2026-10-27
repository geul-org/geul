//ff:func feature=core type=encoder control=sequence
//ff:what EncodeRecord — serialize a Record as TAG WORD, Length extension, Codec extension and Payload
package core

import (
	"errors"
	"fmt"
)

// EncodeRecord serializes r (GEUL t002 README): the TAG WORD, the canonical
// Length extension, the Codec extension TAG (Issuer + format number) when Codec
// is an ISSUER code, then the Payload. An Issuer whose width does not match the
// codec, or extension fields on a direct codec, is a hard error.
func EncodeRecord(r Record) ([]Word, error) {
	if r.Codec > 0x7F {
		return nil, fmt.Errorf("geul/core: codec 0x%02X exceeds 7 bits", r.Codec)
	}
	w := IssuerWidth(r.Kind, r.Codec)
	if w != len(r.Issuer) {
		return nil, fmt.Errorf("geul/core: codec 0x%02X needs a %d-word Issuer, got %d", r.Codec, w, len(r.Issuer))
	}
	if w == 0 && r.Format != 0 {
		return nil, errors.New("geul/core: format number on a direct codec")
	}
	body := append([]Word(nil), r.Issuer...)
	if w > 0 {
		body = append(body, Word(r.Format))
	}
	body = append(body, r.Payload...)
	tag, err := encodeTag(r.Kind, r.Codec, len(body))
	if err != nil {
		return nil, err
	}
	return append(tag, body...), nil
}
