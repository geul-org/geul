//ff:func feature=delta type=encoder control=iteration dimension=1
//ff:what LocalIDs.Define — the Edge Id and LOCAL of a REFERENCE-NODE defining a LOCAL by its Object: a previous pair not yet taken, else two new IDs
package delta

import "github.com/geul-org/gogeul/pkg/core"

// Define returns, for a REFERENCE-NODE that defines a LOCAL by Object
// (Subject LOCAL → object), its Edge Id and the LOCAL: the first pair a
// REFERENCE-NODE with an equal Object had in the previous version that no
// call took yet (so two LOCALs bound to the same Object both carry over),
// else two new IDs. The caller writes REFERENCE-NODE(edge, LOCAL local,
// object) and uses local wherever it means that object.
func (l *LocalIDs) Define(object core.Record) (edge, local uint64) {
	ws, _ := core.EncodeRecord(object)
	for _, p := range l.defs[edgeKey(edgeReferenceNode, ws)] {
		if !l.used[p[0]] && !l.used[p[1]] {
			l.used[p[0]], l.used[p[1]] = true, true
			return p[0], p[1]
		}
	}
	return l.fresh(), l.fresh()
}
