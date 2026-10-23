//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what verdictOf — Fail on any Error, Review on only Warns, Pass on none
package schema

// verdictOf adds Facts up to a Verdict.
func verdictOf(fs []Fact) Verdict {
	v := Pass
	for _, f := range fs {
		if f.Severity == Error {
			return Fail
		}
		v = Review
	}
	return v
}
