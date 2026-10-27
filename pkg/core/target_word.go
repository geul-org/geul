//ff:func feature=core type=helper control=sequence
//ff:what TargetWord — a record's Kind·Codec as a TAG WORD with every other bit 0
package core

// TargetWord is the TAG WORD of a k/codec record with only the Kind and Codec
// bits set (LLLL = 0): the DECLARE target and the codec spec document ID.
func TargetWord(k Kind, codec uint8) Word {
	return Word(k)<<11 | Word(codec&0x7F)
}
