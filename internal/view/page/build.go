//ff:func feature=page type=builder control=sequence
//ff:what Build — the page of an interpreted document
package page

import "github.com/geul-org/gogeul/pkg/inspect"

// Build lays out d under title: HEAD rows, the Edges of the BODY and its nodes
// — the NODE records of a TOULMIN document, else the records inside the Edges
// (refNodes) — with labels and back-references resolved, the SOURCE text and
// the graph: TOULMIN NODE/LINK records, else the Edges in RDF style
// (refGraph).
func Build(d inspect.Document, title string) Page {
	labels := labelMap(d.Body)
	p := Page{Title: title, Profile: d.Profile, Head: headRows(d.Head)}
	p.Nodes, p.Edges = splitRows(bodyRows(d.Body, labels))
	p.Graph = buildGraph(d.Body, labels)
	if !d.Toulmin {
		var at map[string]int
		p.Nodes, at = refNodes(d.Body, labels)
		p.Graph = refGraph(d.Body, p.Nodes, at)
	}
	p.Source, p.HasSource = sourceText(d.Body)
	return p
}
