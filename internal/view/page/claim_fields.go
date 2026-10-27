//ff:func feature=page type=builder control=selection
//ff:what claimFields — a Claim as its subject link, predicate and mode, then its object
package page

import (
	"fmt"
	"strconv"

	"github.com/geul-org/gogeul/pkg/toulmin"
)

// claimFields shows a Claim: the subject link with the predicate and mode
// codes, then the object — a link, an integer or a SOURCE span.
func claimFields(c toulmin.Claim, labels map[uint64]string) []Field {
	claim := Field{Key: "claim", Text: fmt.Sprintf("predicate 0x%04X mode 0x%04X", c.Predicate, c.Mode), Refs: refs([]toulmin.TID{c.Subject}, labels)}
	obj := Field{Key: "object"}
	switch c.Kind {
	case toulmin.ObjNum:
		obj.Text = strconv.FormatUint(uint64(c.ObjNum), 10)
	case toulmin.ObjSpan:
		obj.Text = fmt.Sprintf("@%d+%d", c.ObjSpan.Off, c.ObjSpan.Len)
	default:
		obj.Refs = refs([]toulmin.TID{c.ObjTID}, labels)
	}
	return []Field{claim, obj}
}
