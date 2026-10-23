//ff:func feature=resolve type=helper control=sequence
//ff:what selection.add — record the Node fetched for a Key, reusing an existing Node of the same version
package resolve

// add records n as fetched for k and returns the Node kept. When the same
// version was already fetched (a latest request landing on a pinned version),
// the earlier Node is kept so the graph has one Node per version.
func (s selection) add(k Key, n *Node) *Node {
	nk := NodeKey{DocKey: n.Key.DocKey(), Version: n.Version}
	if cur, ok := s.nodes[nk]; ok {
		n = cur
	}
	s.nodes[nk] = n
	if k.Version == 0 {
		s.latest[k.DocKey()] = n
	}
	return n
}
