//ff:type feature=resolve type=model
//ff:what Graph — a root document and every 판본 it depends on, one Node per (document, version)
package resolve

// Graph is a resolved dependency graph. Docs has one Node per (document,
// version). New versions promise no compatibility, so imports are never merged
// onto one version: each import is bound to the version it asked for.
type Graph struct {
	Root *Node
	Docs map[NodeKey]*Node
}
