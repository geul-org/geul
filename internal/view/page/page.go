//ff:type feature=page type=model
//ff:what Page — the viewer page: title, profile, HEAD and BODY rows and the SOURCE text
package page

// Page is the viewer page of one document. Nodes are the TOULMIN NODE records,
// or in any other document the records inside its Edges; Edges are the BODY's
// Edge records. Nothing else of the BODY is shown. Source is the SOURCE text when the BODY has one
// (HasSource), which the rows' spans point into. Graph is the NODE/LINK graph
// of a TOULMIN document. Browse turns on the tree browser (the 탐색 menu and its
// right sidebar), which only the .well-known/geul tree mode can answer.
type Page struct {
	Title     string
	Profile   string
	Head      []Row
	Nodes     []Row
	Edges     []Row
	Source    string
	HasSource bool
	Graph     Graph
	Browse    bool
}
