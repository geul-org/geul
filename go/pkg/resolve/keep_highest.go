//ff:func feature=resolve type=helper control=sequence
//ff:what keepHighest — record a Node as its document's selection unless a higher version is already selected
package resolve

// keepHighest applies minimal version selection: the highest version asked for
// anywhere in the graph wins (GEUL-ORG 의존성 해석).
func keepHighest(sel map[DocKey]*Node, n *Node) {
	d := n.Key.DocKey()
	if cur, ok := sel[d]; !ok || n.Version > cur.Version {
		sel[d] = n
	}
}
