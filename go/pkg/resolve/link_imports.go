//ff:func feature=resolve type=helper control=iteration dimension=1
//ff:what linkImports — point a Node's IMPORT IDs at the Node each import is bound to
package resolve

// linkImports fills n.Imports from its IMPORT-DOC records with the Node each
// import was bound to: its pinned version, or the latest it resolved to.
func linkImports(n *Node, sel selection) {
	n.Imports = map[uint64]*Node{}
	for id, k := range importKeys(n.Doc) {
		n.Imports[id] = sel.of(k)
	}
}
