//ff:func feature=resolve type=helper control=iteration dimension=1
//ff:what matchAny — whether any pattern selects a document
package resolve

// matchAny reports whether one of ps matches d.
func matchAny(ps []Pattern, d DocKey) bool {
	for _, p := range ps {
		if p.Match(d) {
			return true
		}
	}
	return false
}
