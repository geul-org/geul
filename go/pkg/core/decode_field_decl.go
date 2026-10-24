//ff:func feature=core type=decoder control=selection
//ff:what DecodeFieldDecl — a FIELD Edge as FieldDecl
package core

import "fmt"

// DecodeFieldDecl reads a FIELD Edge (GEUL t002 docs/edge/FIELD): the LOCAL Edge
// Id, the LOCAL Reference to its DECLARE, the width word and, for a record
// field only, an expected TAG word.
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
	f := FieldDecl{EdgeID: id, Declare: decl}
	rest := r.Payload[pos+n:]
	switch {
	case len(rest) == 1:
		f.Width = rest[0]
		return f, nil
	case len(rest) == 2 && rest[0] == 0:
		f.Expect, f.HasExpect = rest[1], true
		return f, nil
	default:
		return FieldDecl{}, fmt.Errorf("geul/core: FIELD needs a width word and, for a record field, one expected TAG (got %d words)", len(rest))
	}
}
