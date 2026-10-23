//ff:func feature=resolve type=helper control=sequence
//ff:what selection.of — the Node a Key is bound to, nil when not fetched yet
package resolve

// of returns the Node k is bound to: the latest Node for version 0, otherwise
// the Node of that exact version.
func (s selection) of(k Key) *Node {
	if k.Version == 0 {
		return s.latest[k.DocKey()]
	}
	return s.nodes[NodeKey{DocKey: k.DocKey(), Version: k.Version}]
}
