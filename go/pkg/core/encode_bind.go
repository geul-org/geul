//ff:func feature=core type=encoder control=sequence
//ff:what EncodeBind — a BIND Edge: LOCAL Edge Id, the Subject Reference, the name as UTF8-TEXT
package core

import (
	"errors"
	"slices"
)

// EncodeBind builds a BIND Edge (docs/edge/BIND) binding b.Name to the
// Subject Reference record b.Subject under Edge Id b.EdgeID (not 0).
func EncodeBind(b Bind) (Record, error) {
	if b.EdgeID == 0 {
		return Record{}, errors.New("geul/core: BIND Edge Id is 0")
	}
	id, _ := EncodeRecord(EncodeLocalRef(b.EdgeID))
	subj, err := EncodeRecord(b.Subject)
	if err != nil {
		return Record{}, err
	}
	name, _ := EncodeRecord(EncodeBytesValue(CodecUTF8TextEven, b.Name))
	return Record{Kind: KindEdge, Codec: EdgeBind, Payload: slices.Concat(id, subj, name)}, nil
}
