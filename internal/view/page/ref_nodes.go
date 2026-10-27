//ff:func feature=page type=builder control=iteration dimension=1
//ff:what refNodes — the nodes of a document: the records inside its Edges, each once, Edge Ids, raw fields and Values left out
package page

import "github.com/geul-org/gogeul/pkg/inspect"

// refNodes lists the records embedded in the Edges of body as nodes, each
// distinct record (by its words) once in order of first appearance, with the
// Edges that hold it as back-references. An Edge's own Edge Id, raw fields
// and Values are left out. A node that a BIND names shows that name as its
// label (bindNames). It also returns each node's row index by its words, as
// nodeIndex keys them.
func refNodes(body []inspect.Entry, labels map[uint64]string) ([]Row, map[string]int) {
	var rows []Row
	at := map[string]int{}
	for _, e := range body {
		for _, r := range edgeNodes(e) {
			i := nodeIndex(&rows, at, r)
			rows[i].Back = append(rows[i].Back, refs([]uint64{e.ID}, labels)...)
		}
	}
	for k, n := range bindNames(body) {
		if i, ok := at[k]; ok {
			rows[i].Label = n
		}
	}
	return rows, at
}
