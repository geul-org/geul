//ff:func feature=core type=helper control=selection
//ff:what CodecMatches — whether a k/codec record fits a CODEC target word: exact, or a wildcard (LLLL 1111) over its zero Codec bits
package core

// CodecMatches reports whether a record of Kind k and Codec codec fits the
// CODEC target t (GEUL t002 docs/edge/CODEC). With LLLL 0000 the Kind and
// Codec must be equal (a Reference's Quantifier included). With LLLL 1111 the
// target is a wildcard: the Kind must be equal and a Codec bit group that is
// all zero matches anything — for a Reference the codec CCCC and the
// Quantifier QQQ are separate groups, for other Kinds the whole 7 bits are one.
func CodecMatches(t Word, k Kind, codec uint8) bool {
	tk, tc := Kind(t>>11&3), uint8(t&0x7F)
	switch {
	case tk != k:
		return false
	case t&0x0780 == 0:
		return tc == codec
	case k != KindReference:
		return tc == 0
	default:
		return (tc>>3 == 0 || tc>>3 == codec>>3) && (tc&7 == 0 || tc&7 == codec&7)
	}
}
