//ff:func feature=toulmin type=encoder control=selection
//ff:what encodeEdgeDst — append a LINK dst as a FuncRef or a REF by the relation's dst kind
package toulmin

import (
	"errors"

	"github.com/geul-org/geul/go/pkg/core"
)

// encodeEdgeDst appends p's dst: a FuncRef for do/undo, a REF for every TID or
// TEXT dst. A FuncRef relation without DstFunc and a provenance block dst (no
// layout yet) are errors.
func encodeEdgeDst(out []core.Word, p EdgePacket, w Wire) ([]core.Word, error) {
	kind, err := PropDstKind(p.Prop)
	if err != nil {
		return nil, err
	}
	switch kind {
	case DstFuncRef:
		if p.DstFunc == nil {
			return nil, errors.New("geul/toulmin: FuncRef dst PropCode but DstFunc is nil")
		}
		return encodeFuncRef(out, *p.DstFunc), nil
	case DstTID, DstText:
		return PutID(out, p.DstTID), nil
	default:
		return nil, errors.New("geul/toulmin: provenance/Block dst has no layout")
	}
}
