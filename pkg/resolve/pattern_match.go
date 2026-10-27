//ff:func feature=resolve type=helper control=sequence
//ff:what Pattern.Match — whether a document falls under the pattern
package resolve

// Match reports whether d falls under p.
func (p Pattern) Match(d DocKey) bool {
	return (p.Issuer == "*" || p.Issuer == d.Issuer) && (p.Doc == "" || p.Doc == d.Doc)
}
