//ff:func feature=pack type=helper control=sequence
//ff:what overlaps — whether entry i of a block (a leaf key, or an internal child's key span) meets [lo, hi)
package pack

// overlaps tests entry i against [lo, hi); hi "" is unbounded. An internal
// child i spans [keys[i], keys[i+1]).
func overlaps(b block, i int, lo, hi string) bool {
	k := b.keys[i]
	if hi != "" && k >= hi {
		return false
	}
	if isLeaf(b.kind) {
		return k >= lo
	}
	return i+1 == len(b.keys) || b.keys[i+1] > lo
}
