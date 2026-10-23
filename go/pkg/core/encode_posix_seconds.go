//ff:func feature=core type=encoder control=sequence
//ff:what EncodePosixSeconds — a POSIX-TIMESTAMP-SECONDS Value: fixed 4-word signed second count
package core

// EncodePosixSeconds builds the POSIX-TIMESTAMP-SECONDS Value for s seconds
// from the Unix epoch, always 4 words (GEUL t002 docs/value/time).
func EncodePosixSeconds(s int64) Record {
	u := uint64(s)
	return Record{Kind: KindValue, Codec: CodecPosixSeconds, Payload: []Word{Word(u >> 48), Word(u >> 32), Word(u >> 16), Word(u)}}
}
