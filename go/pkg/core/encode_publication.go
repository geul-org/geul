//ff:func feature=core type=encoder control=selection
//ff:what EncodePublication — a PUBLICATION, DEPRECATION or TOMBSTONE Edge from a Publication
package core

// EncodePublication builds the version list record p (GEUL t002
// docs/edge/PUBLICATION): the LOCAL Edge Id, the 4-word document ID, the 2-word
// version ID and the codec's tail — the content hash for PUBLICATION, the reason
// for TOMBSTONE, nothing for DEPRECATION.
func EncodePublication(p Publication) Record {
	id, _ := EncodeRecord(EncodeLocalRef(p.EdgeID))
	payload := append(id, p.Doc...)
	payload = append(payload, Word(p.Version>>16), Word(p.Version))
	switch p.Codec {
	case EdgePublication:
		h, _ := ReadWords(p.Hash[:])
		payload = append(payload, h...)
	case EdgeTombstone:
		payload = append(payload, Word(p.Reason))
	}
	return Record{Kind: KindEdge, Codec: p.Codec, Payload: payload}
}
