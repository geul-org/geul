//ff:func feature=page type=helper control=sequence
//ff:what nodeLabel — a graph node's caption: its LABEL name, or its record name and Edge Id
package page

import (
	"strconv"

	"github.com/geul-org/gogeul/pkg/inspect"
)

// nodeLabel captions a graph node with the name a LABEL gives it, or with its
// record name and Edge Id when it has none.
func nodeLabel(e inspect.Entry, labels map[uint64]string) string {
	if l := labels[e.ID]; l != "" {
		return l
	}
	return e.Name + " #" + strconv.FormatUint(e.ID, 10)
}
