//ff:func feature=core type=encoder control=iteration dimension=1
//ff:what intWords — the minimal big-endian two's-complement word form of a signed integer
package core

// intWords returns c as the fewest big-endian two's-complement words: a leading
// 0x0000 or 0xFFFF word appears only when it carries the sign (GEUL t002
// SIGNED-INTEGER minimal-word rule).
func intWords(c int64) []Word {
	out := []Word{Word(uint16(c))}
	for c >>= 16; !(c == 0 && out[0]&0x8000 == 0) && !(c == -1 && out[0]&0x8000 != 0); c >>= 16 {
		out = append([]Word{Word(uint16(c))}, out...)
	}
	return out
}
