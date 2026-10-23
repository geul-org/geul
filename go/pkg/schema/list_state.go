//ff:type feature=schema type=model
//ff:what listState — what a version list has said so far: versions per document, names, the list's Issuer
package schema

// listState is the state of a version list being checked: the version state of
// each document ID, each bound name's document ID, and the list's own Issuer
// (from its DOC-ID) in canonical hex.
type listState struct {
	docs   map[string]*docVersions
	names  map[string]string
	issuer string
}
