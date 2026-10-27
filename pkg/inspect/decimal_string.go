//ff:func feature=inspect type=decoder control=selection
//ff:what decimalString — C×10^E written out exactly as a decimal numeral, keeping its quantum (120×10^-2 is 1.20)
package inspect

import (
	"math/big"
	"strings"
)

// decimalString writes c×10^e as a plain decimal numeral without rounding:
// the digits of c with the point e places from the right ("1.20" for c = 120,
// e = -2; "0.05" for 5, -2; "4200" for 42, 2) and a leading - when negative.
func decimalString(e int32, c *big.Int) string {
	digits := new(big.Int).Abs(c).String()
	sign := ""
	if c.Sign() < 0 {
		sign = "-"
	}
	switch point := len(digits) + int(e); {
	case e >= 0:
		return sign + digits + strings.Repeat("0", int(e))
	case point <= 0:
		return sign + "0." + strings.Repeat("0", -point) + digits
	default:
		return sign + digits[:point] + "." + digits[point:]
	}
}
