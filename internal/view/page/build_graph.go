//ff:func feature=page type=builder control=iteration dimension=1
//ff:what buildGraph — NODE records as nodes and the LINK records joining two of them as edges
package page

import (
	"github.com/geul-org/gogeul/pkg/inspect"
	"github.com/geul-org/gogeul/pkg/toulmin"
)

// buildGraph collects every decoded NODE as a node, labelled by its LABEL
// name, then as edges every decoded LINK whose src and dst (a do/undo's
// function ALIAS) are both nodes, and the NODE-to-NODE references in NODE
// fields. LABEL links and links to TEXT records are left out: they name
// things rather than relate them.
func buildGraph(body []inspect.Entry, labels map[uint64]string) Graph {
	g := Graph{Nodes: []GraphNode{}, Edges: []GraphEdge{}}
	nodes := map[uint64]bool{}
	for _, e := range body {
		if _, ok := e.Data.(toulmin.NodePacket); ok {
			nodes[e.ID] = true
			g.Nodes = append(g.Nodes, GraphNode{ID: anchorOf(e.ID), Label: nodeLabel(e, labels), Kind: e.Name, Class: graphClass(e.Name)})
		}
	}
	for _, e := range body {
		if p, ok := e.Data.(toulmin.NodePacket); ok {
			g.Edges = append(g.Edges, nodeEdges(p, nodes)...)
		}
		if _, ok := e.Data.(toulmin.EdgePacket); ok && len(e.Refs) == 2 && nodes[e.Refs[0]] && nodes[e.Refs[1]] {
			g.Edges = append(g.Edges, GraphEdge{ID: anchorOf(e.ID), Source: anchorOf(e.Refs[0]), Target: anchorOf(e.Refs[1]), Kind: e.Name, Class: graphClass(e.Name)})
		}
	}
	return g
}
