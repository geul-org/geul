//ff:func feature=core type=decoder control=iteration dimension=1
//ff:what ScaleDecimal — rescale C×10^E to an exact integer count of 10^target units
package core

import "errors"

// ScaleDecimal returns c×10^(e−target), the value c×10^e expressed in units of
// 10^target, when that is an exact integer within ±limit. Extra digits below
// 10^target (a non-zero remainder) and magnitudes beyond limit are errors, so a
// caller never silently rounds or overflows.
func ScaleDecimal(e int32, c int64, target int32, limit int64) (int64, error) {
	if c == 0 {
		return 0, nil
	}
	for ; e > target; e-- {
		if c > limit || c < -limit {
			return 0, errors.New("geul/core: decimal value out of range")
		}
		c *= 10
	}
	for ; e < target; e++ {
		if c%10 != 0 {
			return 0, errors.New("geul/core: decimal value has more digits than allowed")
		}
		c /= 10
	}
	if c > limit || c < -limit {
		return 0, errors.New("geul/core: decimal value out of range")
	}
	return c, nil
}
