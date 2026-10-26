//ff:func feature=notation type=decoder control=selection
//ff:what decimalText — a decimal C×10^E as written: digits with a point for E<0, an exponent for E>0
package notation

import (
	"math/big"
	"strconv"
	"strings"
)

// decimalText writes c×10^e keeping its digits (quantum): E 0 is the
// integer, E < 0 places a point (zeros padded: 5×10^-3 is 0.005), E > 0 adds
// `e` and the exponent (12e2), so reading it back gives the same C and E.
func decimalText(e int32, c *big.Int) string {
	digits := new(big.Int).Abs(c).String()
	sign := ""
	if c.Sign() < 0 {
		sign = "-"
	}
	switch {
	case e == 0:
		return sign + digits
	case e > 0:
		return sign + digits + "e" + strconv.Itoa(int(e))
	}
	n := int(-e)
	if len(digits) <= n {
		digits = strings.Repeat("0", n-len(digits)+1) + digits
	}
	return sign + digits[:len(digits)-n] + "." + digits[len(digits)-n:]
}
