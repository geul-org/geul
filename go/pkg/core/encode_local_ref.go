//ff:func feature=core type=encoder control=sequence
//ff:what EncodeLocalRef — a LOCAL Reference record (Quantifier Identified) for an ID in minimal words
package core

// EncodeLocalRef builds the LOCAL Reference record `0000 0LLL L110 0000` + id for
// a document-local ID (GEUL t002 OVERVIEW Reference): Quantifier Identified and
// the target ID in the fewest words. Edge Ids and in-document mentions use it.
func EncodeLocalRef(id uint64) Record {
	return Record{Kind: KindReference, Codec: CodecLocalIdentified, Payload: uintWords(id)}
}
