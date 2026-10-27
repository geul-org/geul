//ff:func feature=toulmin type=decoder control=selection
//ff:what decodeClaimObject — read a Claim object of the given kind at *pos into c
package toulmin

import (
	"errors"
	"fmt"
	"math"

	"github.com/geul-org/gogeul/pkg/core"
)

// decodeClaimObject reads the Claim object of kind at ws[*pos] into c and
// advances *pos: a REF, an UNSIGNED-INTEGER Value (which must fit 32 bits) or a
// SPAN. An unknown kind is an error.
func decodeClaimObject(ws []core.Word, pos *int, kind ObjKind, w Wire, c *Claim) error {
	switch kind {
	case ObjTID:
		t, err := TakeID(ws, pos)
		c.ObjTID = t
		return err
	case ObjNum:
		r, err := TakeValue(ws, pos)
		if err != nil {
			return err
		}
		n, err := core.DecodeUintValue(r)
		if err == nil && n > math.MaxUint32 {
			err = errors.New("geul/toulmin: Claim integer object exceeds 32 bits")
		}
		c.ObjNum = uint32(n)
		return err
	case ObjSpan:
		s, err := TakeSpan(ws, pos, w.Span)
		c.ObjSpan = s
		return err
	default:
		return fmt.Errorf("geul/toulmin: unknown Claim ObjKind %d", uint16(kind))
	}
}
