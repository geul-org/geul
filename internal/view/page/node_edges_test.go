//ff:func feature=page type=builder control=sequence
//ff:what TestNodeEdges — uses and checking edges to graph nodes; references outside the graph draw nothing
package page

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/toulmin"
)

func TestNodeEdges(t *testing.T) {
	c := toulmin.TID(3)
	p := toulmin.NodePacket{TID: 1, Func: &toulmin.FuncRef{Alias: 2}, Check: &c}
	got := nodeEdges(p, map[uint64]bool{2: true, 3: true})
	if len(got) != 2 || got[0] != (GraphEdge{ID: "e1-uses", Source: "e1", Target: "e2", Kind: "uses", Class: "def"}) || got[1].Target != "e3" || got[1].Class != "run" {
		t.Fatalf("got %+v", got)
	}
	if got := nodeEdges(p, map[uint64]bool{}); got != nil {
		t.Fatalf("outside: %+v", got)
	}
}
