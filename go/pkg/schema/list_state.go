//ff:type feature=schema type=model
//ff:what listState — what a version list has said so far: its document, version state and bound names
package schema

// listState is the state of a version list being checked: the document it lists
// (canonical hex), that document's version state and the names bound to it.
type listState struct {
	doc   string
	docs  map[string]*docVersions
	names map[string]string
}
