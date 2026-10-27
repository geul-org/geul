//ff:type feature=page type=model
//ff:what GraphNode — one node in the graph view
package page

// GraphNode is one node: ID is its row anchor (`e<Edge Id>`, `n<i>`) or, for
// a node without a row of its own, an id under its Edge's anchor. Label is its
// name, Kind its record name and Class its style class. Row, when set, is the
// row a tap on it selects.
type GraphNode struct {
	ID    string `json:"id"`
	Label string `json:"label"`
	Kind  string `json:"kind"`
	Class string `json:"class"`
	Row   string `json:"row,omitempty"`
}
