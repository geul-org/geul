//ff:func feature=page type=builder control=sequence
//ff:what TestRefGraph — Edges as hubs, a shared node once, Values as literals, a LOCAL naming an Edge as an ellipse with its ID line; non-Edges no hub
package page

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/inspect"
)

func TestRefGraph(t *testing.T) {
	a, b, c := bindEntry(1, 7, "a"), bindEntry(2, 1, "b"), bindEntry(3, 7, "c")
	a.Name, b.Name, c.Name = "BIND", "BIND", "BIND"
	body := []inspect.Entry{a, b, c, inspectEntry(core.EncodeUintValue(5), nil)}
	rows, at := refNodes(body, nil)
	g := refGraph(body, rows, at)
	if !g.RDF || len(g.Nodes) != 8 || g.Nodes[0] != (GraphNode{ID: "e1", Label: "BIND", Kind: "BIND", Class: "hub"}) {
		t.Fatalf("nodes: %+v", g.Nodes)
	}
	if n := g.Nodes[1]; n.ID != "n0" || n.Label != "c" || n.Class != "res" {
		t.Fatalf("shared node: %+v", n)
	}
	if n := g.Nodes[4]; n.ID != "n1" || n.Label != "b" || n.Class != "res" {
		t.Fatalf("edge node: %+v", n)
	}
	if len(g.Edges) != 7 || g.Edges[2] != (GraphEdge{ID: "e2-2", Source: "e2", Target: "n1", Kind: "BIND", Class: "head", Label: "subject", Row: "e2"}) || g.Edges[4].ID != "n1-id" || g.Edges[4].Target != "e1" {
		t.Fatalf("edges: %+v", g.Edges)
	}
}
