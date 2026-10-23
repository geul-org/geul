//ff:func feature=schema type=helper control=sequence
//ff:what factAt — a Fact at a record
package schema

// factAt builds a Fact located at record r.
func factAt(rule string, sev Severity, r rec, part, expected, actual string) Fact {
	return Fact{Rule: rule, Severity: sev, Where: Where{Word: r.Word, Record: r.Index, Part: part}, Expected: expected, Actual: actual}
}
