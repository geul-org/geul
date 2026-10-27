//ff:func feature=page type=builder control=iteration dimension=1
//ff:what headRows — one row per HEAD record: a Meta as a meta row, a HEAD Edge (document SOURCE, its QUALIFIER) as an Edge row
package page

import (
	"strconv"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/inspect"
)

// headRows lays out each HEAD Meta record as a row anchored `h<n>`, its fields
// summarized for the list (summaryOf). A HEAD Edge (a SOURCE on the whole
// document or a QUALIFIER on one) is laid out like a BODY Edge (bodyRow),
// anchored by its Edge Id so references to it resolve.
func headRows(head []inspect.Entry) []Row {
	rows := make([]Row, 0, len(head))
	for i, e := range head {
		if e.Record.Kind == core.KindEdge {
			rows = append(rows, bodyRow(i, e, nil, nil, 0))
			continue
		}
		r := Row{Index: i, Anchor: "h" + strconv.Itoa(i), Name: rowName(e), Class: "meta", Fields: metaFields(e)}
		r.Summary = summaryOf(r.Fields)
		r.HeadParts, r.BodyParts = hexParts(e)
		rows = append(rows, r)
	}
	return rows
}
