//ff:func feature=core type=decoder control=sequence
//ff:what DecodePublication — a PUBLICATION, DEPRECATION or TOMBSTONE Edge as a Publication
package core

import (
	"errors"
	"fmt"
)

// DecodePublication reads a version list record (GEUL t002 docs/edge/PUBLICATION):
// the LOCAL Edge Id, the 4-word document ID, the 2-word version ID (not 0) and
// the codec's tail. The length after the Edge Id is fixed per codec.
func DecodePublication(r Record) (Publication, error) {
	tail, ok := publicationTail(r.Codec)
	if r.Kind != KindEdge || !ok {
		return Publication{}, fmt.Errorf("geul/core: not a version list record (kind %d codec 0x%02X)", r.Kind, r.Codec)
	}
	ref, n, err := DecodeRecord(r.Payload)
	if err != nil {
		return Publication{}, err
	}
	id, err := DecodeLocalRef(ref)
	if err != nil {
		return Publication{}, err
	}
	rest := r.Payload[n:]
	if len(rest) != 6+tail {
		return Publication{}, fmt.Errorf("geul/core: version list record needs %d words after the Edge Id, got %d", 6+tail, len(rest))
	}
	p := Publication{Codec: r.Codec, EdgeID: id, Doc: rest[:4], Version: uint32(rest[4])<<16 | uint32(rest[5])}
	if p.Version == 0 {
		return Publication{}, errors.New("geul/core: version ID 0 in a version list record")
	}
	fillPublicationTail(&p, rest[6:])
	return p, nil
}
