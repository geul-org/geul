//ff:func feature=page type=helper control=iteration dimension=1
//ff:what addEdges — graph lines appended once each by id
package page

// addEdges appends to edges each of es whose id seen has not had yet, and
// marks it seen.
func addEdges(edges []GraphEdge, seen map[string]bool, es []GraphEdge) []GraphEdge {
	for _, l := range es {
		if !seen[l.ID] {
			seen[l.ID] = true
			edges = append(edges, l)
		}
	}
	return edges
}
