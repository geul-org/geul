//ff:func feature=toulmin type=encoder control=sequence
//ff:what encodeNodeSlots — append a NODE's flag-gated fields in flag order, rejecting a set flag without its field
package toulmin

import (
	"errors"

	"github.com/geul-org/geul/go/pkg/core"
)

// encodeNodeSlots appends p's conditional fields in flag order (F_FUNC, F_QUAL,
// F_CHECK, F_CLAIM, F_VERDICT, F_WITH, F_GROUNDS). A flag whose field is absent is an error
// so a record never carries a truncated field.
func encodeNodeSlots(out []core.Word, p NodePacket, w Wire) ([]core.Word, error) {
	var err error
	if p.Flags&NodeFlagFunc != 0 {
		if p.Func == nil {
			return nil, errors.New("geul/toulmin: F_FUNC set but Func is nil")
		}
		out = encodeFuncRef(out, *p.Func)
	}
	if err == nil && p.Flags&NodeFlagQual != 0 {
		if p.Qual == nil {
			return nil, errors.New("geul/toulmin: F_QUAL set but Qual is nil")
		}
		out, err = PutValue(out, EncodeVFixValue(*p.Qual))
	}
	if err == nil && p.Flags&NodeFlagCheck != 0 {
		if p.Check == nil {
			return nil, errors.New("geul/toulmin: F_CHECK set but Check is nil")
		}
		out = PutID(out, *p.Check)
	}
	if err == nil && p.Flags&NodeFlagClaim != 0 {
		if p.Claim == nil {
			return nil, errors.New("geul/toulmin: F_CLAIM set but Claim is nil")
		}
		out, err = encodeClaim(out, *p.Claim, w)
	}
	if err == nil && p.Flags&NodeFlagVerdict != 0 {
		if p.Verdict == nil {
			return nil, errors.New("geul/toulmin: F_VERDICT set but Verdict is nil")
		}
		out, err = PutValue(out, EncodeVFixValue(*p.Verdict))
	}
	if err != nil {
		return nil, err
	}
	if p.Flags&NodeFlagWith != 0 {
		out = encodeWith(out, p.With)
	}
	if p.Flags&NodeFlagGrounds != 0 {
		out = encodeWith(out, p.Grounds)
	}
	return out, nil
}
