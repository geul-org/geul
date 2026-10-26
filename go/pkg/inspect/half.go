//ff:func feature=inspect type=decoder control=selection
//ff:what half — an IEEE 754 binary16 bit pattern as float64: normal, subnormal, ±0, ±Inf and NaN
package inspect

import "math"

// half decodes the binary16 bits h exactly (every binary16 is a float64).
func half(h uint16) float64 {
	sign := 1.0
	if h&0x8000 != 0 {
		sign = -1
	}
	exp, frac := int(h>>10&0x1F), float64(h&0x3FF)
	switch {
	case exp == 0x1F && frac == 0:
		return math.Inf(int(sign))
	case exp == 0x1F:
		return math.NaN()
	case exp == 0:
		return sign * math.Ldexp(frac, -24)
	default:
		return sign * math.Ldexp(1024+frac, exp-25)
	}
}
