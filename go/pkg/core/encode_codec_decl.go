//ff:func feature=core type=encoder control=sequence
//ff:what EncodeCodecDecl — a CODEC Edge from a CodecDecl
package core

// EncodeCodecDecl builds the CODEC Edge c (GEUL t002 docs/edge/CODEC).
func EncodeCodecDecl(c CodecDecl) Record {
	p := append(localWords(c.EdgeID), localWords(c.Field)...)
	return Record{Kind: KindEdge, Codec: EdgeCodec, Payload: append(p, c.Target)}
}
