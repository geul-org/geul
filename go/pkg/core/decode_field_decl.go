//ff:func feature=core type=decoder control=selection
//ff:what DecodeFieldDecl — a FIELD Edge as FieldDecl
package core

import "fmt"

// DecodeFieldDecl reads a FIELD Edge (GEUL t002 docs/edge/FIELD): the LOCAL Edge
// Id, the width word and, for a record field only, an expected TAG word.
func DecodeFieldDecl(r Record) (FieldDecl, error) {
	if r.Kind != KindEdge || r.Codec != EdgeField {
		return FieldDecl{}, fmt.Errorf("geul/core: not a FIELD Edge (kind %d codec 0x%02X)", r.Kind, r.Codec)
	}
	id, pos, err := takeLocal(r.Payload)
	if err != nil {
		return FieldDecl{}, err
	}
	rest := r.Payload[pos:]
	switch {
	case len(rest) == 1:
		return FieldDecl{EdgeID: id, Width: rest[0]}, nil
	case len(rest) == 2 && rest[0] == 0:
		return FieldDecl{EdgeID: id, Expect: rest[1], HasExpect: true}, nil
	default:
		return FieldDecl{}, fmt.Errorf("geul/core: FIELD needs a width word and, for a record field, one expected TAG (got %d words)", len(rest))
	}
}
