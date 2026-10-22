//ff:func feature=geulmd type=encoder control=selection
//ff:what claimStr — a Claim as `#subject/0xpred/<object>/0xmode`
package geulmd

import (
	"fmt"

	"github.com/geul-org/geul/go/pkg/toulmin"
)

// claimStr renders a Claim: subject REF, predicate code, the object (a REF
// `#n`, an integer `n:5` or a span `@o+l`) and mode code, slash-separated.
func claimStr(c toulmin.Claim) string {
	var obj string
	switch c.Kind {
	case toulmin.ObjNum:
		obj = fmt.Sprintf("n:%d", c.ObjNum)
	case toulmin.ObjSpan:
		obj = spanStr(c.ObjSpan)
	default:
		obj = fmt.Sprintf("#%d", c.ObjTID)
	}
	return fmt.Sprintf("#%d/0x%04x/%s/0x%04x", c.Subject, c.Predicate, obj, c.Mode)
}
