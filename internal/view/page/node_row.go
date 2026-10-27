//ff:func feature=page type=builder control=sequence
//ff:what nodeRow — the row of one node record found inside Edges
package page

import (
	"strconv"

	"github.com/geul-org/gogeul/pkg/inspect"
)

// nodeRow lays out node record e as the i-th node, anchored `n<i>`: its codec
// name, its fields in hex as the label (refLabel), the record read as text
// for the list (recordText: Q42, #3, …), and its words.
func nodeRow(i int, e inspect.Entry) Row {
	r := Row{Index: i, Anchor: "n" + strconv.Itoa(i), Name: e.Name, Class: "ref"}
	r.Label = refLabel(e.Record)
	r.Summary = recordText(e.Record)
	r.HeadParts, r.BodyParts = hexParts(e)
	return r
}
