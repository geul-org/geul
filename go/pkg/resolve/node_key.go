//ff:type feature=resolve type=model
//ff:what NodeKey — one 판본: a document identity plus the concrete version fetched
package resolve

// NodeKey identifies one fetched version (판본) of a document. A graph keeps one
// Node per NodeKey, so several versions of the same document can sit side by
// side (OVERVIEW 계보 개정).
type NodeKey struct {
	DocKey
	Version uint32
}
