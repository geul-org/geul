//ff:func feature=toulmin type=decoder control=selection
//ff:what decodeEdgeDst — read a LINK dst (FuncRef or REF) at *pos into p
package toulmin

import (
	"errors"

	"github.com/geul-org/gogeul/pkg/core"
)

// decodeEdgeDst reads p's dst at ws[*pos] in the form p.Prop dictates.
func decodeEdgeDst(ws []core.Word, pos *int, p *EdgePacket, w Wire) error {
	kind, err := PropDstKind(p.Prop)
	if err != nil {
		return err
	}
	switch kind {
	case DstFuncRef:
		f, err := decodeFuncRef(ws, pos, w)
		p.DstFunc = &f
		return err
	case DstTID, DstText:
		p.DstTID, err = TakeID(ws, pos)
		return err
	default:
		return errors.New("geul/toulmin: provenance/Block dst has no layout")
	}
}
