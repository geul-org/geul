//ff:func feature=geulmd type=encoder control=sequence
//ff:what linkLine — a LINK record as `<NAME> #id #src -> <dst> [flags] [gate] @span`
package geulmd

import (
	"fmt"

	"github.com/geul-org/geul/go/pkg/toulmin"
)

// gateOps names the confidence-gate operators.
var gateOps = [4]string{"at-least", "above", "less-than", "at-most"}

// linkLine renders p, a decoded LINK: its src and dst (a FuncRef for do/undo),
// the flags DO-ONCE does not already imply (once on UNDO, strike, ordered), the
// gate as `gate=<op>/<percent>`, then the span.
func linkLine(name string, p toulmin.EdgePacket) string {
	dst := fmt.Sprintf("#%d", p.DstTID)
	if p.DstFunc != nil {
		dst = funcStr(*p.DstFunc)
	}
	s := fmt.Sprintf("%s #%d #%d -> %s", name, p.ID, p.Src, dst)
	if p.Flags&toulmin.EdgeFlagOnce != 0 && name != "DO-ONCE" {
		s += " once"
	}
	if p.Flags&toulmin.EdgeFlagStrike != 0 {
		s += " strike"
	}
	if p.Flags&toulmin.EdgeFlagOrdered != 0 {
		s += " ordered"
	}
	if p.Gate != nil {
		s += fmt.Sprintf(" gate=%s/%d", gateOps[p.Gate.Op&3], p.Gate.Percent)
	}
	return s + " " + spanStr(p.Span)
}
