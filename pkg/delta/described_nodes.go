//ff:func feature=delta type=decoder control=iteration dimension=1
//ff:what describedNodes — the LOCAL nodes of a previous BODY grouped by what their TRIPLEs say of them, in first-use order
package delta

import "github.com/geul-org/gogeul/pkg/core"

// describedNodes groups the TRIPLEs of body by their LOCAL Subject
// (tripleAbout) in document order and maps the key of each Subject's list
// (nodeKey) to the Subjects that have it, first used first, for Node.
func describedNodes(body []core.Record) map[string][]uint64 {
	about := map[uint64][][]core.Word{}
	var order []uint64
	for _, r := range body {
		id, words, ok := tripleAbout(r)
		if !ok {
			continue
		}
		if _, seen := about[id]; !seen {
			order = append(order, id)
		}
		about[id] = append(about[id], words)
	}
	nodes := map[string][]uint64{}
	for _, id := range order {
		k := nodeKey(about[id])
		nodes[k] = append(nodes[k], id)
	}
	return nodes
}
