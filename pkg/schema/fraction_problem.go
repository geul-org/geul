//ff:func feature=schema type=rule control=sequence
//ff:what fractionProblem — why a FRACTION's components are invalid: non-canonical numerator or denominator, a DECIMAL with exponent 0, a zero denominator; "" when valid
package schema

import "github.com/geul-org/gogeul/pkg/core"

// fractionProblem applies the FRACTION component rules (value/numeric/FRACTION)
// to f, whose codecs core.DecodeFraction has checked: each component canonical
// (numberProblem), no DECIMAL component with exponent 0 (one notation, one
// byte sequence), and a denominator whose value is not 0.
func fractionProblem(f core.Fraction) string {
	if p := numberProblem(f.Numerator); p != "" {
		return "numerator: " + p
	}
	if p := numberProblem(f.Denominator); p != "" {
		return "denominator: " + p
	}
	if zeroExpDecimal(f.Numerator) || zeroExpDecimal(f.Denominator) {
		return "지수 0인 DECIMAL 성분(정수 원문은 INTEGER)"
	}
	if d, err := core.NumberRat(f.Denominator); err != nil || d.Sign() == 0 {
		return "분모 0"
	}
	return ""
}
