//ff:func feature=inspect type=decoder control=iteration dimension=1
//ff:what TestDecimalString — quantum kept, leading zeros below one, trailing zeros for positive exponents, sign and 31 digits
package inspect

import (
	"math/big"
	"testing"
)

func TestDecimalString(t *testing.T) {
	big31, _ := new(big.Int).SetString("-1234567890123456789012345678901", 10)
	for want, c := range map[string]struct {
		e int32
		c *big.Int
	}{
		"1.20": {-2, big.NewInt(120)}, "0.05": {-2, big.NewInt(5)}, "4200": {2, big.NewInt(42)},
		"-157.249": {-3, big.NewInt(-157249)}, "652": {0, big.NewInt(652)}, "-123456789012345678901234.5678901": {-7, big31},
	} {
		if got := decimalString(c.e, c.c); got != want {
			t.Fatalf("got %q want %q", got, want)
		}
	}
}
