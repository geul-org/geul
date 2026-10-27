//ff:type feature=page type=model
//ff:what Graph — the graph view: TOULMIN NODE/LINK records, or any document's Edges drawn in RDF style
package page

// Graph is the graph view's data, written into the page as JSON. RDF marks a
// graph drawn from Edges (refGraph) rather than TOULMIN NODE/LINK records.
type Graph struct {
	RDF   bool        `json:"rdf"`
	Nodes []GraphNode `json:"nodes"`
	Edges []GraphEdge `json:"edges"`
}
