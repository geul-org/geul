//ff:type feature=resolve type=model
//ff:what selection — fetched Nodes by (document, version) and the Node each latest request resolved to
package resolve

// selection is what collect has fetched: nodes by (document, version), and for
// latest (version 0) requests the Node each document resolved to in this run.
type selection struct {
	nodes  map[NodeKey]*Node
	latest map[DocKey]*Node
}
