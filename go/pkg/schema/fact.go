//ff:type feature=schema type=model
//ff:what Fact — one finding: rule ID, severity, place, expected and actual
package schema

// Fact is one finding. Rule is the rule ID of the README tables ("F3", "R2" …).
type Fact struct {
	Rule     string
	Severity Severity
	Where    Where
	Expected string
	Actual   string
}
