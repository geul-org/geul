//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what headSourcesOf — the Edge Ids of a draft HEAD's SOURCEs on the whole document
package cli

import "github.com/geul-org/gogeul/pkg/core"

// headSourcesOf returns the Edge Id of every SOURCE in head that is a HEAD
// Edge (core.HeadEdge: Target LOCAL 0).
func headSourcesOf(head []core.Record) map[uint64]bool {
	out := map[uint64]bool{}
	for _, r := range head {
		if id, ok := core.HeadEdge(r, out); ok && r.Codec == core.EdgeSource {
			out[id] = true
		}
	}
	return out
}
