//ff:func feature=pack type=helper control=sequence
//ff:what defaultInt — v when at least lo, def otherwise
package pack

// defaultInt keeps v ≥ lo and replaces anything smaller with def.
func defaultInt(v, lo, def int) int {
	if v < lo {
		return def
	}
	return v
}
