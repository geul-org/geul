//ff:func feature=core type=decoder control=sequence
//ff:what DecodeBind — a BIND Edge as Bind: LOCAL Edge Id, one Reference, one UTF8-TEXT Value
package core

import (
	"errors"
	"fmt"
)

// DecodeBind reads a BIND Edge (GEUL t002 docs/edge/BIND): the LOCAL Edge Id,
// exactly one Reference record (the Subject) and exactly one UTF8-TEXT Value
// (the name), with nothing after it.
func DecodeBind(r Record) (Bind, error) {
	if r.Kind != KindEdge || r.Codec != EdgeBind {
		return Bind{}, fmt.Errorf("geul/core: not a BIND Edge (kind %d codec 0x%02X)", r.Kind, r.Codec)
	}
	ref, n, err := DecodeRecord(r.Payload)
	if err != nil {
		return Bind{}, err
	}
	id, err := DecodeLocalRef(ref)
	if err != nil {
		return Bind{}, err
	}
	subj, m, err := DecodeRecord(r.Payload[n:])
	if err != nil || subj.Kind != KindReference {
		return Bind{}, errors.New("geul/core: BIND Subject must be one Reference record")
	}
	name, k, err := DecodeRecord(r.Payload[n+m:])
	if err != nil || n+m+k != len(r.Payload) {
		return Bind{}, errors.New("geul/core: BIND needs one UTF8-TEXT name ending the payload")
	}
	b, err := DecodeBytesValue(name, CodecUTF8TextEven)
	if err != nil {
		return Bind{}, err
	}
	return Bind{EdgeID: id, Subject: subj, Name: b}, nil
}
