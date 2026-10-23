//ff:type feature=resolve type=model
//ff:what Graph — a root document and every document it depends on, one Node per document
package resolve

// Graph is a resolved dependency graph. Docs has one Node per document: every
// import of the same document, whatever version it pinned, shares the Node of
// the selected version.
type Graph struct {
	Root *Node
	Docs map[DocKey]*Node
}
