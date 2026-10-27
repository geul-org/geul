//ff:func feature=page type=builder control=sequence
//ff:what bodyRow — the row of one BODY record
package page

import (
	"strconv"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/inspect"
)

// bodyRow lays out e: anchored by its Edge Id (`e<id>`) or its position
// (`r<n>`), with its label, fields and their one-line summary (summaryOf),
// back-references, span and words.
func bodyRow(i int, e inspect.Entry, labels map[uint64]string, back []uint64, depth int) Row {
	r := Row{Index: i, Anchor: "r" + strconv.Itoa(i), Name: rowName(e), Class: classOf(e.Data), ID: e.ID,
		Label: labels[e.ID], Fields: bodyFields(e, labels), Back: refs(back, labels), Depth: depth, Edge: e.Record.Kind == core.KindEdge}
	if e.ID != 0 {
		r.Anchor = anchorOf(e.ID)
	}
	r.Summary = summaryOf(r.Fields)
	sp, ok := spanOf(e.Data)
	r.HasSpan, r.SpanOff, r.SpanLen = ok, sp.Off, sp.Len
	r.HeadParts, r.BodyParts = hexParts(e)
	return r
}
