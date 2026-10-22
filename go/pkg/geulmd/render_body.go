//ff:func feature=geulmd type=encoder control=iteration dimension=1
//ff:what renderBody — the BODY block lines: structured where lossless, raw otherwise
package geulmd

import (
	"slices"

	"github.com/geul-org/geul/go/pkg/core"
	"github.com/geul-org/geul/go/pkg/toulmin"
)

// renderBody renders each BODY record as one line in order. A standalone Value
// shows its literal. Under the TANGEUL profile a TOULMIN record shows its
// structured form — the SOURCE first sets the SPAN width — but only when
// re-encoding that form reproduces the record's words exactly; every other
// record, and any record that fails that check, is a raw RECORD line.
func renderBody(body []core.Record, ctx headCtx) []string {
	w := toulmin.Wire{ID: ctx.idWidth, Span: 1}
	lines := make([]string, 0, len(body))
	for _, r := range body {
		orig, _ := core.EncodeRecord(r)
		line, enc, err := structuredLine(r, ctx, &w)
		if err != nil || !slices.Equal(enc, orig) {
			line = rawRecord(r)
		}
		lines = append(lines, line)
	}
	return lines
}
