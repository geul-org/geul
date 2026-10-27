//ff:func feature=page type=helper control=iteration dimension=1
//ff:what headEnd — which end of an Edge its arrow points at: the object (second node), else the subject (only node)
package page

// headEnd returns the index in ends of the node the Edge's arrow points at,
// RDF style: with two or more nodes (ellipses) the second, the object, so the
// subject runs through the hub into it; with one node that node, so an Edge
// with only a subject points at it; with none -1.
func headEnd(ends []GraphNode) int {
	head, count := -1, 0
	for i, n := range ends {
		if n.Class == "res" && count < 2 {
			head, count = i, count+1
		}
	}
	return head
}
