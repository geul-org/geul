//ff:func feature=page type=builder control=sequence
//ff:what TestBuild — HEAD rows, TOULMIN nodes or nodes inside Edges, Edge rows, labels, back-references, case nesting, the SOURCE text and the TOULMIN or RDF graph
package page

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/inspect"
)

func TestBuild(t *testing.T) {
	p := Build(fixtureDoc(t), "x.geul")
	if p.Title != "x.geul" || p.Profile != "tangeul-1.0" || len(p.Head) != 2 || len(p.Nodes) != 2 || len(p.Edges) != 8 || !p.HasSource || p.Source != "# 사안\n규칙" {
		t.Fatalf("got %+v", p)
	}
	cs := p.Nodes[0]
	if cs.Anchor != "e1" || cs.Label != "사안" || cs.Class != "node" || len(cs.Back) != 5 || !cs.HasSpan || cs.SpanOff != 2 {
		t.Fatalf("case row: %+v", cs)
	}
	if p.Edges[4].Depth != 0 || p.Edges[5].Depth != 1 || p.Edges[6].Depth != 0 {
		t.Fatalf("depths: %d %d %d", p.Edges[4].Depth, p.Edges[5].Depth, p.Edges[6].Depth)
	}
	if len(p.Graph.Nodes) != 2 || len(p.Graph.Edges) != 2 || p.Graph.Nodes[0].Label != "사안" {
		t.Fatalf("graph: %+v", p.Graph)
	}
	if raw := p.Edges[7]; raw.Class != "raw" || raw.Anchor != "e16" || raw.Name != "PREDICATE-DECLARE" {
		t.Fatalf("raw row: %+v", raw)
	}
	plain := inspect.Document{Body: []inspect.Entry{bindEntry(1, 7, "x")}}
	if p := Build(plain, "y"); len(p.Nodes) != 1 || p.Nodes[0].Name != "LOCAL" || len(p.Edges) != 1 || !p.Graph.RDF || len(p.Graph.Edges) != 2 {
		t.Fatalf("plain: %+v", p)
	}
}
