//ff:func feature=core type=encoder control=sequence
//ff:what EncodeNextLocalID — a NEXT-LOCAL-ID Meta: the next LOCAL ID to assign as a minimal unsigned integer
package core

// EncodeNextLocalID builds NEXT-LOCAL-ID (meta/lineage/NEXT-LOCAL-ID) for
// next, the LOCAL ID the next new statement gets, in the fewest words.
func EncodeNextLocalID(next uint64) Record {
	return Record{Kind: KindMeta, Codec: MetaNextLocalID, Payload: uintWords(next)}
}
