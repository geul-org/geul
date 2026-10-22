//ff:func feature=toulmin type=decoder control=selection
//ff:what decodeControlArg — read a control's argument (SPAN, section code or case REF) at *pos into p
package toulmin

import (
	"github.com/geul-org/geul/go/pkg/core"
)

// decodeControlArg reads the argument of p.Code at ws[*pos]: the SPAN of a
// PROSE, the 1-word section code of a SECTION, or the case REF of CASE-BEGIN /
// CASE-END.
func decodeControlArg(ws []core.Word, pos *int, p *ControlPacket, w Wire) error {
	switch p.Code {
	case ControlProse:
		s, err := TakeSpan(ws, pos, w.Span)
		p.Span = &s
		return err
	case ControlSection:
		sec, err := TakeWord(ws, pos, "SECTION code")
		p.Arg = uint64(sec)
		return err
	default:
		t, err := TakeID(ws, pos, w.ID)
		p.Arg = uint64(t)
		return err
	}
}
