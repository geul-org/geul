//ff:func feature=delta type=decoder control=iteration dimension=1
//ff:what addedEdges — the Edges of a snapshot whose Edge Id the other version did not have, in order
package delta

// addedEdges lists n's Edges whose IDs are not in before as '+'.
func addedEdges(n snapshot, before map[uint64]bool) []Change {
	var out []Change
	for i, id := range n.ids {
		if !before[id] {
			out = append(out, Change{Op: '+', ID: id, New: n.body[i]})
		}
	}
	return out
}
