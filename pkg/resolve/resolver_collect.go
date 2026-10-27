//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what Resolver.collect — fetch every 판본 reachable from the roots, one Node per (document, version)
package resolve

import "context"

// collect walks the import graph breadth first. Each import is bound to the
// version it asks for; nothing is merged onto a common version. A Key already
// bound is not fetched again, and cycles end because bound Keys are skipped.
func (r *Resolver) collect(ctx context.Context, roots []Key) (selection, error) {
	sel := selection{nodes: map[NodeKey]*Node{}, latest: map[DocKey]*Node{}}
	queue := roots
	for len(queue) > 0 {
		k := queue[0]
		queue = queue[1:]
		if sel.of(k) != nil {
			continue
		}
		n, err := r.resolveKey(ctx, k)
		if err != nil {
			return selection{}, err
		}
		n = sel.add(k, n)
		for _, ik := range importKeys(n.Doc) {
			queue = append(queue, ik)
		}
	}
	return sel, nil
}
