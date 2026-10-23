//ff:type feature=resolve type=model
//ff:what Node — one fetched, verified document in a Graph with its resolved imports
package resolve

import "github.com/geul-org/geul/go/pkg/inspect"

// Node is one resolved document. Key is what was asked for; Version and Hash
// are what was fetched and verified. Imports maps this document's IMPORT IDs to
// the Nodes the graph selected for them. Source names the GEULPROXY item it
// came from. The root Node of a Graph has an empty Key and Source.
type Node struct {
	Key     Key
	Version uint32
	Hash    [32]byte
	Doc     inspect.Document
	Imports map[uint64]*Node
	Source  string
}
