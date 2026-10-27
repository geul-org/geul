//ff:func feature=schema type=rule control=sequence
//ff:what checkFraction — V10: a FRACTION holds a SIGNED-INTEGER/DECIMAL numerator and an UNSIGNED-INTEGER/DECIMAL denominator, canonical, no exponent-0 DECIMAL, non-zero denominator, and nothing more
package schema

import "github.com/geul-org/gogeul/pkg/core"

// checkFraction applies V10 (value/numeric/FRACTION): the two records decode
// in their forms (core.DecodeFraction) and their components are valid
// (fractionProblem).
func checkFraction(r rec) []Fact {
	f, err := core.DecodeFraction(r.R)
	if err != nil {
		return []Fact{factAt("V10", Error, r, "Payload", "numerator·denominator 레코드", err.Error())}
	}
	if p := fractionProblem(f); p != "" {
		return []Fact{factAt("V10", Error, r, "Payload", "정규 성분, 지수 0 아닌 DECIMAL, 0 아닌 분모", p)}
	}
	return nil
}
