//ff:func feature=page type=builder control=sequence
//ff:what TestBuildGraph — NODEs become labelled nodes; LINKs between nodes and NODE field references become edges; LABEL and TEXT links do not
package page

import "testing"

func TestBuildGraph(t *testing.T) {
	g := buildGraph(fixtureDoc(t).Body, map[uint64]string{1: "사안"})
	if len(g.Nodes) != 2 || g.Nodes[0] != (GraphNode{ID: "e1", Label: "사안", Kind: "CASE-DECLARE", Class: "case"}) || g.Nodes[1].Label != "RULE-DECLARE #2" {
		t.Fatalf("nodes: %+v", g.Nodes)
	}
	if len(g.Edges) != 2 || g.Edges[0].ID != "e2-checking" || g.Edges[1] != (GraphEdge{ID: "e14", Source: "e2", Target: "e1", Kind: "ATTACKS", Class: "attack"}) {
		t.Fatalf("edges: %+v", g.Edges)
	}
}
