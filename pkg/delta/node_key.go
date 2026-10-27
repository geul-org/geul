//ff:func feature=delta type=helper control=sequence
//ff:what nodeKey — the content key of a described node: the words of what its TRIPLEs say, in order
package delta

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// nodeKey identifies a described node by its TRIPLEs (Node, describedNodes):
// two nodes whose lists of Predicate·Object·Rank words match are the same node.
func nodeKey(about [][]core.Word) string {
	return fmt.Sprintf("%04x", about)
}
