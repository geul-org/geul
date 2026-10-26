//ff:func feature=notation type=decoder control=iteration dimension=1
//ff:what TestDecimalText — integers, points with padding, exponents and signs keep every digit
package notation

import (
	"math/big"
	"testing"
)

func TestDecimalText(t *testing.T) {
	cases := []struct {
		e    int32
		c    int64
		want string
	}{{0, 652, "652"}, {-2, 120, "1.20"}, {-3, 5, "0.005"}, {2, 12, "12e2"}, {-1, -15, "-1.5"}, {0, -3, "-3"}, {-2, 0, "0.00"}}
	for _, c := range cases {
		if got := decimalText(c.e, big.NewInt(c.c)); got != c.want {
			t.Errorf("%d×10^%d: %s, want %s", c.c, c.e, got, c.want)
		}
	}
}
