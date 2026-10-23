//ff:func feature=resolve type=helper control=iteration dimension=1
//ff:what linkImports — point a Node's IMPORT IDs at the selected Node of each imported document
package resolve

// linkImports fills n.Imports from its IMPORT-DOC-* records with the Nodes sel
// selected for those documents.
func linkImports(n *Node, sel map[DocKey]*Node) {
	n.Imports = map[uint64]*Node{}
	for id, k := range importKeys(n.Doc) {
		n.Imports[id] = sel[k.DocKey()]
	}
}
