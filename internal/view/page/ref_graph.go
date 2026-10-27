//ff:func feature=page type=builder control=iteration dimension=1
//ff:what refGraph — a document's Edges in RDF style: Edges as diamond hubs, nodes as ellipses, Values as boxes
package page

import (
	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/inspect"
)

// refGraph draws every Edge of body as a hub labelled by its codec name, then
// its lines (edgeGraph). A node several Edges share is drawn once, and so is
// the idLine of a LOCAL several Edges hold.
func refGraph(body []inspect.Entry, rows []Row, at map[string]int) Graph {
	g := Graph{RDF: true, Nodes: []GraphNode{}, Edges: []GraphEdge{}}
	hubs := map[uint64]bool{}
	for _, e := range body {
		hubs[e.ID] = e.ID != 0 && e.Record.Kind == core.KindEdge
	}
	seen := map[string]bool{}
	for _, e := range body {
		if !hubs[e.ID] {
			continue
		}
		g.Nodes = append(g.Nodes, GraphNode{ID: anchorOf(e.ID), Label: e.Name, Kind: e.Name, Class: "hub"})
		ns, es := edgeGraph(e, hubs, rows, at)
		g.Nodes = addNodes(g.Nodes, seen, ns)
		g.Edges = addEdges(g.Edges, seen, es)
	}
	return g
}
