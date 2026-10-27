//ff:func feature=page type=helper control=sequence
//ff:what nodeIndex — the row index of a node record, adding its row the first time its words are seen
package page

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/inspect"
)

// nodeIndex returns the index in rows of node record r, keyed by its words in
// at, appending a new row when r is seen for the first time.
func nodeIndex(rows *[]Row, at map[string]int, r inspect.Entry) int {
	k := fmt.Sprint(r.Words)
	if i, seen := at[k]; seen {
		return i
	}
	at[k] = len(*rows)
	*rows = append(*rows, nodeRow(len(*rows), r))
	return at[k]
}
