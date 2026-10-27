//ff:func feature=page type=builder control=iteration dimension=1
//ff:what headRows — one row per HEAD Meta record
package page

import (
	"strconv"

	"github.com/geul-org/gogeul/pkg/inspect"
)

// headRows lays out each HEAD Meta record as a row anchored `h<n>`, its fields
// summarized for the list (summaryOf).
func headRows(head []inspect.Entry) []Row {
	rows := make([]Row, 0, len(head))
	for i, e := range head {
		r := Row{Index: i, Anchor: "h" + strconv.Itoa(i), Name: rowName(e), Class: "meta", Fields: metaFields(e)}
		r.Summary = summaryOf(r.Fields)
		r.HeadParts, r.BodyParts = hexParts(e)
		rows = append(rows, r)
	}
	return rows
}
