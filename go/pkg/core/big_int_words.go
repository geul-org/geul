//ff:func feature=core type=encoder control=iteration dimension=1
//ff:what bigIntWords — an arbitrary-size signed integer as minimal big-endian two's complement WORDs
package core

import "math/big"

// bigIntWords is intWords for any size: c in the fewest big-endian two's
// complement WORDs (a leading sign WORD only when the next WORD's top bit
// would read as the wrong sign).
func bigIntWords(c *big.Int) []Word {
	mask := big.NewInt(0xFFFF)
	v := new(big.Int).Set(c)
	low := new(big.Int)
	out := []Word{Word(low.And(v, mask).Uint64())}
	for v.Rsh(v, 16); !(v.Sign() == 0 && out[0]&0x8000 == 0) && !(v.Cmp(big.NewInt(-1)) == 0 && out[0]&0x8000 != 0); v.Rsh(v, 16) {
		out = append([]Word{Word(low.And(v, mask).Uint64())}, out...)
	}
	return out
}
