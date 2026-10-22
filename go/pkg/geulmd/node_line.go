//ff:func feature=geulmd type=encoder control=sequence
//ff:what nodeLine — a NODE record as `<NAME> #id role= level= [fields…] @span`
package geulmd

import (
	"fmt"
	"strconv"

	"github.com/geul-org/geul/go/pkg/toulmin"
)

// nodeLine renders p, a decoded NODE, with every field its flags carry so the
// line holds the whole record: role and level codes, then each present field in
// flag order (func, qual, check, claim, verdict, with, grounds, strike), then
// the span.
func nodeLine(name string, p toulmin.NodePacket) string {
	s := fmt.Sprintf("%s #%d role=%d level=%d", name, p.TID, p.Role, p.Level)
	if p.Func != nil {
		s += " func=" + funcStr(*p.Func)
	}
	if p.Qual != nil {
		s += " qual=" + strconv.FormatFloat(p.Qual.Float(), 'f', -1, 64)
	}
	if p.Check != nil {
		s += fmt.Sprintf(" check=#%d", *p.Check)
	}
	if p.Claim != nil {
		s += " claim=" + claimStr(*p.Claim)
	}
	if p.Verdict != nil {
		s += " verdict=" + strconv.FormatFloat(p.Verdict.Float(), 'f', -1, 64)
	}
	if p.Flags&toulmin.NodeFlagWith != 0 {
		s += " with=" + tidList(p.With)
	}
	if p.Flags&toulmin.NodeFlagGrounds != 0 {
		s += " grounds=" + tidList(p.Grounds)
	}
	if p.Flags&toulmin.NodeFlagStrike != 0 {
		s += " strike"
	}
	return s + " " + spanStr(p.Span)
}
