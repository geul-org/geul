//ff:func feature=notation type=encoder control=sequence
//ff:what parseDecimal — C and E of a decimal written as digits, an optional point and an optional exponent, digits kept
package notation

import (
	"fmt"
	"math/big"
	"regexp"
	"strconv"
)

// decimalPattern is sign, integer digits, fraction digits, exponent.
var decimalPattern = regexp.MustCompile(`^([+-]?)([0-9]*)(?:\.([0-9]+))?(?:[eE]([+-]?[0-9]+))?$`)

// parseDecimal reads s as c×10^e keeping every written digit: `1.20` is
// 120×10^-2, `12e2` is 12×10^2 (decimalText reads back).
func parseDecimal(s string) (int32, *big.Int, error) {
	m := decimalPattern.FindStringSubmatch(s)
	if m == nil || m[2]+m[3] == "" {
		return 0, nil, fmt.Errorf("geul/notation: %q is not a decimal number", s)
	}
	c, _ := new(big.Int).SetString(m[1]+m[2]+m[3], 10)
	exp, err := int64(0), error(nil)
	if m[4] != "" {
		exp, err = strconv.ParseInt(m[4], 10, 32)
	}
	e := exp - int64(len(m[3]))
	if err != nil || e < -1<<31 || e > 1<<31-1 {
		return 0, nil, fmt.Errorf("geul/notation: exponent of %q out of range", s)
	}
	return int32(e), c, nil
}
