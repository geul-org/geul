//ff:func feature=page type=builder control=sequence
//ff:what nodeEdges — the edges a NODE's own fields draw: its function's ALIAS (uses) and its checked case
package page

import "github.com/geul-org/gogeul/pkg/toulmin"

// nodeEdges draws the NODE-to-NODE references a NODE carries in its fields:
// F_FUNC's ALIAS as a `uses` edge and F_CHECK's case as a `checking` edge,
// each only when the target is a graph node. Their ids add a suffix to the
// NODE's anchor, as they are not records of their own.
func nodeEdges(p toulmin.NodePacket, nodes map[uint64]bool) []GraphEdge {
	var out []GraphEdge
	from := anchorOf(uint64(p.TID))
	if p.Func != nil && nodes[uint64(p.Func.Alias)] {
		out = append(out, GraphEdge{ID: from + "-uses", Source: from, Target: anchorOf(uint64(p.Func.Alias)), Kind: "uses", Class: "def"})
	}
	if p.Check != nil && nodes[uint64(*p.Check)] {
		out = append(out, GraphEdge{ID: from + "-checking", Source: from, Target: anchorOf(uint64(*p.Check)), Kind: "checking", Class: "run"})
	}
	return out
}
