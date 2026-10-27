//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what Resolver.Resolve — decode a root document and bind its IMPORT-DOC dependencies into a Graph
package resolve

import (
	"context"

	"github.com/geul-org/gogeul/pkg/inspect"
)

// Resolve builds the dependency graph of root: every 판본 reachable through
// IMPORT-DOC, each import bound to its own version, one Node per (document,
// version). Any document that cannot be fetched or verified fails the whole
// resolution.
func (r *Resolver) Resolve(ctx context.Context, root []byte) (*Graph, error) {
	doc, err := inspect.Decode(root)
	if err != nil {
		return nil, err
	}
	g := &Graph{Root: &Node{Doc: doc}}
	var keys []Key
	for _, k := range importKeys(doc) {
		keys = append(keys, k)
	}
	sel, err := r.collect(ctx, keys)
	if err != nil {
		return nil, err
	}
	g.Docs = sel.nodes
	linkImports(g.Root, sel)
	for _, n := range g.Docs {
		linkImports(n, sel)
	}
	return g, nil
}
