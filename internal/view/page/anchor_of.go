//ff:func feature=page type=helper control=sequence
//ff:what anchorOf — the row anchor of a record with an Edge Id (`e<id>`)
package page

import "strconv"

// anchorOf is the row anchor of the record with Edge Id id; the graph view
// uses the same ids, so a node and its row point at each other.
func anchorOf(id uint64) string {
	return "e" + strconv.FormatUint(id, 10)
}
