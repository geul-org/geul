//ff:func feature=core type=encoder control=sequence
//ff:what EncodeFieldDecl — a FIELD Edge from a FieldDecl
package core

// EncodeFieldDecl builds the FIELD Edge f (GEUL t002 docs/edge/FIELD).
func EncodeFieldDecl(f FieldDecl) Record {
	p := append(append(localWords(f.EdgeID), localWords(f.Declare)...), f.Width)
	return Record{Kind: KindEdge, Codec: EdgeField, Payload: p}
}
