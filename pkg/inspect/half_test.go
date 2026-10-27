//ff:func feature=inspect type=decoder control=iteration dimension=1
//ff:what TestHalf — normal, subnormal, signed zero, infinities and NaN binary16 values
package inspect

import (
	"math"
	"testing"
)

func TestHalf(t *testing.T) {
	for h, want := range map[uint16]float64{0x3C00: 1, 0xC000: -2, 0x7BFF: 65504, 0x0001: math.Ldexp(1, -24), 0x0000: 0, 0x7C00: math.Inf(1), 0xFC00: math.Inf(-1)} {
		if got := half(h); got != want {
			t.Fatalf("%04X: %v want %v", h, got, want)
		}
	}
	if !math.Signbit(half(0x8000)) || !math.IsNaN(half(0x7E00)) {
		t.Fatal("-0 or NaN")
	}
}
