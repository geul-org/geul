//ff:func feature=core type=helper control=sequence
//ff:what decimalRat — C×10^E as an exact rational number
package core

import "math/big"

// decimalRat returns c×10^e exactly (no rounding): c over 10^-e for e < 0,
// c times 10^e otherwise.
func decimalRat(e int32, c *big.Int) *big.Rat {
	if e < 0 {
		return new(big.Rat).SetFrac(c, new(big.Int).Exp(big.NewInt(10), big.NewInt(-int64(e)), nil))
	}
	return new(big.Rat).SetInt(new(big.Int).Mul(c, new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(e)), nil)))
}
