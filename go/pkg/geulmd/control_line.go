//ff:func feature=geulmd type=encoder control=selection
//ff:what controlLine — a framing record as SECTION n, CASE-BEGIN/END #case or PROSE @span
package geulmd

import (
	"fmt"

	"github.com/geul-org/geul/go/pkg/toulmin"
)

// controlLine renders a decoded framing record: SECTION with its section code,
// CASE-BEGIN/CASE-END with the case reference, PROSE with its span.
func controlLine(name string, p toulmin.ControlPacket) string {
	switch p.Code {
	case toulmin.ControlProse:
		return fmt.Sprintf("%s #%d %s", name, p.ID, spanStr(*p.Span))
	case toulmin.ControlSection:
		return fmt.Sprintf("%s #%d %d", name, p.ID, p.Arg)
	default:
		return fmt.Sprintf("%s #%d #%d", name, p.ID, p.Arg)
	}
}
