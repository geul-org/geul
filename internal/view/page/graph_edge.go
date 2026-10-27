//ff:type feature=page type=model
//ff:what GraphEdge — one line between two graph nodes
package page

// GraphEdge is one line from Source to Target (node ids). ID is its own id —
// a LINK's row anchor, or an id under its Edge's anchor — Kind its record name
// and Class its style class. Label names the line (an RDF graph's field name)
// and Row, when set, is the row a tap on it selects.
type GraphEdge struct {
	ID     string `json:"id"`
	Source string `json:"source"`
	Target string `json:"target"`
	Kind   string `json:"kind"`
	Class  string `json:"class"`
	Label  string `json:"label,omitempty"`
	Row    string `json:"row,omitempty"`
}
