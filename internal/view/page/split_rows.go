//ff:func feature=page type=builder control=iteration dimension=1
//ff:what splitRows — BODY rows split into TOULMIN NODE rows and Edge rows, each in document order; anything else is not shown
package page

// splitRows splits BODY rows into TOULMIN NODE records and the other Edge
// records, each in document order. Records of another Kind (a Value or a
// Reference standing alone in the BODY) are left out.
func splitRows(rows []Row) ([]Row, []Row) {
	var nodes, edges []Row
	for _, r := range rows {
		if r.Class == "node" {
			nodes = append(nodes, r)
		} else if r.Edge {
			edges = append(edges, r)
		}
	}
	return nodes, edges
}
