//ff:func feature=toulmin type=encoder control=sequence
//ff:what EncodeEdge — encode an EdgePacket as a LINK record (do+once as DO-ONCE)
package toulmin

import (
	"errors"

	"github.com/geul-org/geul/go/pkg/core"
)

// EncodeEdge encodes p as a LINK record (GEUL t002 TOULMIN LINK): Edge Id, src
// REF, dst, flag word, the gate when F_GATE, then the SPAN. A `do` with F_ONCE
// is written under DO-ONCE and the flag is dropped from the wire, since the
// codec already says once.
func EncodeEdge(p EdgePacket, w Wire) ([]core.Word, error) {
	once := p.Prop == PropDo && p.Flags&EdgeFlagOnce != 0
	ref, err := PropCodec(p.Prop, once)
	if err != nil {
		return nil, err
	}
	flags := p.Flags
	if once {
		flags &^= EdgeFlagOnce
	}
	body, err := encodeEdgeDst(PutID(PutID(nil, p.ID), p.Src), p, w)
	if err != nil {
		return nil, err
	}
	body = append(body, core.Word(flags))
	if p.Flags&EdgeFlagGate != 0 {
		if p.Gate == nil {
			return nil, errors.New("geul/toulmin: F_GATE set but Gate is nil")
		}
		g, err := EncodeGate(*p.Gate)
		if err != nil {
			return nil, err
		}
		body = append(body, g...)
	}
	if body, err = PutSpan(body, p.Span, w.Span); err != nil {
		return nil, err
	}
	return EdgeRecord(ref, body)
}
