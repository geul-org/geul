//ff:func feature=toulmin type=decoder control=sequence
//ff:what decodeNodeSlots — read a NODE's flag-gated fields in flag order into p
package toulmin

import (
	"github.com/geul-org/geul/go/pkg/core"
)

// decodeNodeSlots reads p's conditional fields at ws[*pos] in flag order,
// following p.Flags; objKind (from the header) selects the Claim object form.
func decodeNodeSlots(ws []core.Word, pos *int, p *NodePacket, objKind ObjKind, w Wire) error {
	var err error
	if p.Flags&NodeFlagFunc != 0 {
		var f FuncRef
		f, err = decodeFuncRef(ws, pos, w)
		p.Func = &f
	}
	if err == nil && p.Flags&NodeFlagQual != 0 {
		p.Qual, err = TakeVFix(ws, pos)
	}
	if err == nil && p.Flags&NodeFlagCheck != 0 {
		var t TID
		t, err = TakeID(ws, pos, w.ID)
		p.Check = &t
	}
	if err == nil && p.Flags&NodeFlagClaim != 0 {
		var c Claim
		c, err = decodeClaim(ws, pos, objKind, w)
		p.Claim = &c
	}
	if err == nil && p.Flags&NodeFlagVerdict != 0 {
		p.Verdict, err = TakeVFix(ws, pos)
	}
	if err == nil && p.Flags&NodeFlagWith != 0 {
		p.With, err = decodeWith(ws, pos, w)
	}
	if err == nil && p.Flags&NodeFlagGrounds != 0 {
		p.Grounds, err = decodeWith(ws, pos, w)
	}
	return err
}
