//ff:func feature=core type=helper control=selection
//ff:what validTarget — whether a word is a CODEC target: top 3 bits 0 and LLLL 0000, or a 1111 wildcard that leaves some Codec bit group zero
package core

// validTarget reports whether t can be a CODEC target (edge/CODEC): the top 3
// bits are 0 and LLLL is 0000 (exact) or 1111 (wildcard). A wildcard must
// leave something to match: a Value, Edge or Meta wildcard has Codec 0, a
// Reference wildcard has its codec CCCC or its Quantifier QQQ 0.
func validTarget(t Word) bool {
	c := uint8(t & 0x7F)
	switch {
	case t&0xE000 != 0:
		return false
	case t&0x0780 == 0:
		return true
	case t&0x0780 != 0x0780:
		return false
	case Kind(t>>11&3) != KindReference:
		return c == 0
	default:
		return c>>3 == 0 || c&7 == 0
	}
}
