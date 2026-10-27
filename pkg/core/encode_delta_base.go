//ff:func feature=core type=encoder control=sequence
//ff:what EncodeDeltaBase — a DELTA-BASE Meta: the 2-word version ID a diff applies to
package core

// EncodeDeltaBase builds DELTA-BASE (meta/lineage/DELTA-BASE): this version
// file is a diff against version of the same document.
func EncodeDeltaBase(version uint32) Record {
	return Record{Kind: KindMeta, Codec: MetaDeltaBase, Payload: []Word{Word(version >> 16), Word(version)}}
}
