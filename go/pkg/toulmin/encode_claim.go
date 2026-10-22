//ff:func feature=toulmin type=encoder control=selection
//ff:what encodeClaim — append a Claim block: subject REF, predicate, object by kind, mode
package toulmin

import (
	"fmt"

	"github.com/geul-org/geul/go/pkg/core"
)

// encodeClaim appends c to out: the subject REF, the predicate word, the object
// (a REF, an UNSIGNED-INTEGER Value or a SPAN by c.Kind) and the mode word. The
// kind itself goes into the NODE header word. An unknown Kind is an error.
func encodeClaim(out []core.Word, c Claim, w Wire) ([]core.Word, error) {
	out, err := PutID(out, c.Subject, w.ID)
	if err != nil {
		return nil, err
	}
	out = append(out, core.Word(c.Predicate))
	switch c.Kind {
	case ObjTID:
		out, err = PutID(out, c.ObjTID, w.ID)
	case ObjNum:
		out, err = PutValue(out, core.EncodeUintValue(uint64(c.ObjNum)))
	case ObjSpan:
		out, err = PutSpan(out, c.ObjSpan, w.Span)
	default:
		return nil, fmt.Errorf("geul/toulmin: unknown Claim ObjKind %d", uint16(c.Kind))
	}
	if err != nil {
		return nil, err
	}
	return append(out, core.Word(c.Mode)), nil
}
