//ff:func feature=core type=decoder control=sequence
//ff:what DecodeCodecDecl — a CODEC Edge as CodecDecl
package core

import "fmt"

// DecodeCodecDecl reads a CODEC Edge (GEUL t002 docs/edge/CODEC): the LOCAL
// Edge Id, a LOCAL Reference to its FIELD and one target word (validTarget).
func DecodeCodecDecl(r Record) (CodecDecl, error) {
	if r.Kind != KindEdge || r.Codec != EdgeCodec {
		return CodecDecl{}, fmt.Errorf("geul/core: not a CODEC Edge (kind %d codec 0x%02X)", r.Kind, r.Codec)
	}
	id, pos, err := takeLocal(r.Payload)
	if err != nil {
		return CodecDecl{}, err
	}
	f, n, err := takeLocal(r.Payload[pos:])
	if err != nil {
		return CodecDecl{}, err
	}
	if rest := r.Payload[pos+n:]; len(rest) != 1 || !validTarget(rest[0]) {
		return CodecDecl{}, fmt.Errorf("geul/core: CODEC needs one target word, LLLL 0000 or a 1111 wildcard (got %04X)", rest)
	}
	return CodecDecl{EdgeID: id, Field: f, Target: r.Payload[pos+n]}, nil
}
