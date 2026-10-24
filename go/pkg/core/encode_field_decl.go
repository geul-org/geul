//ff:func feature=core type=encoder control=sequence
//ff:what EncodeFieldDecl — a FIELD Edge from a FieldDecl
package core

// EncodeFieldDecl builds the FIELD Edge f (GEUL t002 docs/edge/FIELD). An
// expected TAG is written only for a record field (Width 0).
func EncodeFieldDecl(f FieldDecl) Record {
	p := append(append(localWords(f.EdgeID), localWords(f.Declare)...), f.Width)
	if f.Width == 0 && f.HasExpect {
		p = append(p, f.Expect)
	}
	return Record{Kind: KindEdge, Codec: EdgeField, Payload: p}
}
