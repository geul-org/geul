//ff:func feature=core type=decoder control=sequence
//ff:what DecodeFieldDecl — a FIELD Edge as FieldDecl
package core

import "fmt"

// DecodeFieldDecl reads a FIELD Edge (GEUL t002 docs/edge/FIELD): the LOCAL Edge
// Id, the LOCAL Reference to its DECLARE and the width word.
func DecodeFieldDecl(r Record) (FieldDecl, error) {
	if r.Kind != KindEdge || r.Codec != EdgeField {
		return FieldDecl{}, fmt.Errorf("geul/core: not a FIELD Edge (kind %d codec 0x%02X)", r.Kind, r.Codec)
	}
	id, pos, err := takeLocal(r.Payload)
	if err != nil {
		return FieldDecl{}, err
	}
	decl, n, err := takeLocal(r.Payload[pos:])
	if err != nil {
		return FieldDecl{}, err
	}
	if rest := r.Payload[pos+n:]; len(rest) != 1 {
		return FieldDecl{}, fmt.Errorf("geul/core: FIELD needs one width word (got %d words)", len(rest))
	}
	return FieldDecl{EdgeID: id, Declare: decl, Width: r.Payload[pos+n]}, nil
}
