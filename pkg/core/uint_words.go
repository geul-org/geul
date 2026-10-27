//ff:func feature=core type=encoder control=iteration dimension=1
//ff:what uintWords — the minimal big-endian word form of an unsigned integer (0 is one 0x0000 word)
package core

// uintWords returns n as the fewest big-endian words: 0 is a single 0x0000 word
// and no other value has a leading 0x0000 word (GEUL t002 UNSIGNED-INTEGER).
func uintWords(n uint64) []Word {
	out := []Word{Word(n)}
	for n >>= 16; n != 0; n >>= 16 {
		out = append([]Word{Word(n)}, out...)
	}
	return out
}
