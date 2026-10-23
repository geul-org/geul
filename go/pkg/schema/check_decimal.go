//ff:func feature=schema type=rule control=sequence
//ff:what checkDecimal — V3: exponent 2 words and a coefficient in the fewest words
package schema

// checkDecimal applies V3 to a SIGNED-/UNSIGNED-DECIMAL Value.
func checkDecimal(r rec, signed bool) []Fact {
	if len(r.R.Payload) < 3 {
		return []Fact{factAt("V3", Error, r, "Payload", "exponent 2 WORD + coefficient 1 WORD 이상", fmtWords(len(r.R.Payload)))}
	}
	return checkInteger(r, r.R.Payload[2:], signed, "coefficient")
}
