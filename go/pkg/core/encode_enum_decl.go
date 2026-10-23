//ff:func feature=core type=encoder control=sequence
//ff:what EncodeEnumDecl — an ENUM Edge from an EnumDecl
package core

// EncodeEnumDecl builds the ENUM Edge e (GEUL t002 docs/edge/ENUM).
func EncodeEnumDecl(e EnumDecl) Record {
	p := append(localWords(e.EdgeID), localWords(e.Field)...)
	return Record{Kind: KindEdge, Codec: EdgeEnum, Payload: append(p, e.Code...)}
}
