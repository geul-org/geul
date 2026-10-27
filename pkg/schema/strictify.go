//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what strictify — turn every Warn into an Error (Options.Strict)
package schema

// strictify returns fs with every Warn raised to Error.
func strictify(fs []Fact) []Fact {
	out := make([]Fact, len(fs))
	for i, f := range fs {
		f.Severity = Error
		out[i] = f
	}
	return out
}
