//ff:func feature=delta type=encoder control=iteration dimension=1
//ff:what LocalIDs.Node — the LOCAL of a node known by what TRIPLEs say of it: the previous LOCAL with the same description not yet taken, else a new ID
package delta

import "github.com/geul-org/gogeul/pkg/core"

// Node returns the LOCAL for a node that the document describes only by
// TRIPLEs whose Subject it is (such as a cited source, edge/SOURCE). about
// lists, in document order, each such TRIPLE's words after its Subject's
// place is taken out: Predicate, then Object and Rank (tripleAbout). A node
// of the previous version with the same list keeps its LOCAL; otherwise a new
// ID is handed out. The caller then writes the TRIPLEs with that LOCAL.
func (l *LocalIDs) Node(about [][]core.Word) uint64 {
	for _, id := range l.nodes[nodeKey(about)] {
		if !l.used[id] {
			l.used[id] = true
			return id
		}
	}
	return l.fresh()
}
