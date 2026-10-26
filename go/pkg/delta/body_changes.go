//ff:func feature=delta type=decoder control=iteration dimension=1
//ff:what bodyChanges — Edges of two snapshots matched by Edge Id: removed and changed in base order, then added in next order
package delta

import "github.com/geul-org/geul/go/pkg/core"

// bodyChanges compares the BODY Edges of b and n by Edge Id.
func bodyChanges(b, n snapshot) []Change {
	after := map[uint64]core.Record{}
	for i, id := range n.ids {
		after[id] = n.body[i]
	}
	var out []Change
	before := map[uint64]bool{}
	for i, id := range b.ids {
		before[id] = true
		r, ok := after[id]
		switch {
		case !ok:
			out = append(out, Change{Op: '-', ID: id, Old: b.body[i]})
		case !sameRecord(r, b.body[i]):
			out = append(out, Change{Op: '~', ID: id, Old: b.body[i], New: r})
		}
	}
	return append(out, addedEdges(n, before)...)
}
