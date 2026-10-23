//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what Resolver.Resolve — decode a root document and resolve its IMPORT-DOC-* dependencies into a Graph
package resolve

import (
	"context"

	"github.com/geul-org/geul/go/pkg/inspect"
)

// Resolve builds the dependency graph of root: every document reachable through
// IMPORT-DOC-*, one selected version per document, imports linked to those
// selections. Any document that cannot be fetched or verified fails the whole
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
	if g.Docs, err = r.collect(ctx, keys); err != nil {
		return nil, err
	}
	linkImports(g.Root, g.Docs)
	for _, n := range g.Docs {
		linkImports(n, g.Docs)
	}
	return g, nil
}
