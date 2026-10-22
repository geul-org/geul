//ff:func feature=toulmin type=decoder control=sequence
//ff:what DecodeEdge — decode a LINK record: relation from the codec, src, dst, flags, gate and SPAN
package toulmin

import (
	"errors"
	"fmt"

	"github.com/geul-org/geul/go/pkg/core"
)

// DecodeEdge decodes r as a LINK record (GEUL t002 TOULMIN LINK). The codec gives
// the relation; DO-ONCE restores F_ONCE and must not also carry it on the wire.
// Flag bits above b3 are reserved, and the fields must use the whole payload.
func DecodeEdge(r core.Record, w Wire) (EdgePacket, error) {
	ref, err := RecordRef(r)
	if err != nil {
		return EdgePacket{}, err
	}
	prop, once, ok := PropFromCodec(ref)
	if !ok {
		return EdgePacket{}, fmt.Errorf("geul/toulmin: record 0x%02X/0x%04X is not a LINK", ref.Codec, ref.Format)
	}
	ws, pos := r.Payload, 0
	p := EdgePacket{Prop: prop}
	if p.ID, err = TakeID(ws, &pos, w.ID); err == nil {
		p.Src, err = TakeID(ws, &pos, w.ID)
	}
	if err == nil {
		err = decodeEdgeDst(ws, &pos, &p, w)
	}
	if err != nil {
		return EdgePacket{}, err
	}
	flags, err := TakeWord(ws, &pos, "LINK flag")
	if err != nil {
		return EdgePacket{}, err
	}
	p.Flags = EdgeFlags(flags)
	if flags&0xFFF0 != 0 || once && p.Flags&EdgeFlagOnce != 0 {
		return EdgePacket{}, errors.New("geul/toulmin: LINK reserved flag bits set or F_ONCE on DO-ONCE")
	}
	if once {
		p.Flags |= EdgeFlagOnce
	}
	if p.Flags&EdgeFlagGate != 0 {
		g, n, err := DecodeGate(ws[pos:])
		if err != nil {
			return EdgePacket{}, err
		}
		p.Gate, pos = &g, pos+n
	}
	if p.Span, err = TakeSpan(ws, &pos, w.Span); err != nil {
		return EdgePacket{}, err
	}
	return p, ExpectEnd(ws, pos, "LINK")
}
