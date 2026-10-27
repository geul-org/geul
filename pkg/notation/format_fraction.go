//ff:func feature=notation type=decoder control=sequence
//ff:what formatFraction — FRACTION(numerator, denominator) with each component as written (2, 3 or 23.45, 1256.1)
package notation

import "github.com/geul-org/gogeul/pkg/core"

// formatFraction writes a FRACTION Value (core.DecodeFraction) as
// FRACTION(n, d), each component by fractionPartText. A FRACTION whose
// components have no such text is false (raw words instead).
func formatFraction(r core.Record) (string, bool) {
	f, err := core.DecodeFraction(r)
	if err != nil {
		return "", false
	}
	n, ok1 := fractionPartText(f.Numerator)
	d, ok2 := fractionPartText(f.Denominator)
	if !ok1 || !ok2 {
		return "", false
	}
	return "FRACTION(" + n + ", " + d + ")", true
}
