//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what Resolver.collect — fetch every Key reachable from the roots and keep the highest version per document
package resolve

import (
	"context"
	"fmt"
)

// collect walks the requirement graph breadth first. Each distinct Key is
// fetched once, the imports of every fetched version are followed (as Go's
// minimal version selection reads every go.mod it meets), and the highest
// version of each document is selected. Cycles end because Keys repeat.
func (r *Resolver) collect(ctx context.Context, roots []Key) (map[DocKey]*Node, error) {
	sel := map[DocKey]*Node{}
	seen := map[string]bool{}
	queue := roots
	for len(queue) > 0 {
		k := queue[0]
		queue = queue[1:]
		id := fmt.Sprintf("%v@%d", k.DocKey(), k.Version)
		if seen[id] {
			continue
		}
		seen[id] = true
		n, err := r.Fetch(ctx, k)
		if err != nil {
			return nil, err
		}
		for _, ik := range importKeys(n.Doc) {
			queue = append(queue, ik)
		}
		keepHighest(sel, n)
	}
	return sel, nil
}
