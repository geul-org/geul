//ff:func feature=core type=encoder control=iteration dimension=1
//ff:what bigUintWords — a non-negative integer of any size as minimal big-endian unsigned WORDs (0 is one 0x0000 WORD)
package core

import "math/big"

// bigUintWords is uintWords for any size: c (≥ 0) in the fewest big-endian
// WORDs, 0 as a single 0x0000 WORD.
func bigUintWords(c *big.Int) []Word {
	mask := big.NewInt(0xFFFF)
	v := new(big.Int).Set(c)
	low := new(big.Int)
	out := []Word{Word(low.And(v, mask).Uint64())}
	for v.Rsh(v, 16); v.Sign() != 0; v.Rsh(v, 16) {
		out = append([]Word{Word(low.And(v, mask).Uint64())}, out...)
	}
	return out
}
