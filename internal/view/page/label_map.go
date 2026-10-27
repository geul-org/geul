//ff:func feature=page type=builder control=iteration dimension=1
//ff:what labelMap — Edge Id → name, from the LABEL records that point a record at a TEXT
package page

import (
	"github.com/geul-org/gogeul/pkg/inspect"
	"github.com/geul-org/gogeul/pkg/toulmin"
)

// labelMap returns each labelled record's name: a LABEL record (src → dst)
// names src with the text of the TEXT record dst.
func labelMap(body []inspect.Entry) map[uint64]string {
	texts := map[toulmin.TID]string{}
	for _, e := range body {
		if t, ok := e.Data.(toulmin.TextPacket); ok {
			texts[t.TID] = string(t.Bytes)
		}
	}
	labels := map[uint64]string{}
	for _, e := range body {
		if l, ok := e.Data.(toulmin.EdgePacket); ok && l.Prop == toulmin.PropLabel {
			labels[uint64(l.Src)] = texts[l.DstTID]
		}
	}
	return labels
}
