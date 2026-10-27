//ff:func feature=page type=helper control=sequence
//ff:what graphEnd — the graph node one record field of an Edge stands for: a literal box or a Node row's ellipse
package page

import (
	"fmt"
	"strconv"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/inspect"
)

// graphEnd returns the node that record field p (the k-th part of Edge e)
// stands for. A Value is a literal of its own under e's anchor; any other
// record, a LOCAL naming an Edge included, is the ellipse of its Node row
// (rows, at from refNodes), labelled by the row's name.
func graphEnd(e inspect.Entry, k int, p inspect.Part, rows []Row, at map[string]int) GraphNode {
	r, _, _ := core.DecodeRecord(p.Words)
	if r.Kind == core.KindValue {
		lit := anchorOf(e.ID) + "-" + strconv.Itoa(k) + "v"
		return GraphNode{ID: lit, Label: partText(p), Kind: core.CodecName(r.Kind, r.Codec), Class: "lit", Row: anchorOf(e.ID)}
	}
	row := rows[at[fmt.Sprint(p.Words)]]
	label := row.Label
	if label == "" {
		label = row.Name
	}
	return GraphNode{ID: row.Anchor, Label: label, Kind: row.Name, Class: "res"}
}
