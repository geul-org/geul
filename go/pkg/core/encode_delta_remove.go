//ff:func feature=core type=encoder control=iteration dimension=1
//ff:what EncodeDeltaRemove — a DELTA-REMOVE Meta: one LOCAL Reference per removed LOCAL ID
package core

// EncodeDeltaRemove builds DELTA-REMOVE (meta/lineage/DELTA-REMOVE) for the
// LOCAL IDs a diff removes from its base, in the order given.
func EncodeDeltaRemove(ids []uint64) Record {
	var p []Word
	for _, id := range ids {
		ws, _ := EncodeRecord(EncodeLocalRef(id))
		p = append(p, ws...)
	}
	return Record{Kind: KindMeta, Codec: MetaDeltaRemove, Payload: p}
}
