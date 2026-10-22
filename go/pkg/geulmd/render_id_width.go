//ff:func feature=geulmd type=encoder control=sequence
//ff:what renderIDWidth — `ID-WIDTH edge=E var=V import=I` and the widths it declares
package geulmd

import (
	"fmt"

	"github.com/geul-org/geul/go/pkg/core"
)

// renderIDWidth renders Meta ID-WIDTH as its three widths in words and records
// the Edge Id and IMPORT ID widths in ctx. ok is false for a payload that is not
// one word with the 10 reserved bits clear.
func renderIDWidth(r core.Record, ctx *headCtx) (string, bool) {
	if len(r.Payload) != 1 || r.Payload[0]&0x03FF != 0 {
		return "", false
	}
	w := r.Payload[0]
	e, v, i := core.WidthWords(uint8(w>>14)), core.WidthWords(uint8(w>>12)), core.WidthWords(uint8(w>>10))
	ctx.idWidth, ctx.importWidth = e, i
	return fmt.Sprintf("ID-WIDTH edge=%d var=%d import=%d", e, v, i), true
}
