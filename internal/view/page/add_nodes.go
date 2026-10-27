//ff:func feature=page type=helper control=iteration dimension=1
//ff:what addNodes — graph nodes appended once each by id
package page

// addNodes appends to nodes each of ns whose id seen has not had yet, and
// marks it seen.
func addNodes(nodes []GraphNode, seen map[string]bool, ns []GraphNode) []GraphNode {
	for _, n := range ns {
		if !seen[n.ID] {
			seen[n.ID] = true
			nodes = append(nodes, n)
		}
	}
	return nodes
}
