//ff:func feature=core type=encoder control=sequence
//ff:what EncodeUintValue — an UNSIGNED-INTEGER Value record in minimal words
package core

// EncodeUintValue builds the UNSIGNED-INTEGER Value record (codec 0x02) for n in
// minimal big-endian words.
func EncodeUintValue(n uint64) Record {
	return Record{Kind: KindValue, Codec: CodecUnsignedInteger, Payload: uintWords(n)}
}
