//ff:func feature=page type=builder control=iteration dimension=1
//ff:what bodyRows — one row per BODY record, with back-references and case nesting
package page

import (
	"github.com/geul-org/gogeul/pkg/inspect"
	"github.com/geul-org/gogeul/pkg/toulmin"
)

// bodyRows lays out each BODY record as a row. Records between CASE-BEGIN and
// CASE-END are nested one level deeper; the markers themselves stay outside.
func bodyRows(body []inspect.Entry, labels map[uint64]string) []Row {
	back := backrefMap(body)
	rows := make([]Row, 0, len(body))
	depth := 0
	for i, e := range body {
		c, ctl := e.Data.(toulmin.ControlPacket)
		if ctl && c.Code == toulmin.ControlCaseEnd && depth > 0 {
			depth--
		}
		rows = append(rows, bodyRow(i, e, labels, back[e.ID], depth))
		if ctl && c.Code == toulmin.ControlCaseBegin {
			depth++
		}
	}
	return rows
}
