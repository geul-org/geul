//ff:func feature=page type=builder control=sequence
//ff:what linkFields — a LINK's src and dst links, its flags, its gate and its attack force
package page

import (
	"fmt"
	"strconv"

	"github.com/geul-org/gogeul/pkg/toulmin"
)

// gateOps names the confidence-gate operators (GEUL t002 TOULMIN LINK).
var gateOps = [4]string{"≥", ">", "<", "≤"}

// linkFields lays out a LINK: src, dst (a FuncRef for do/undo), the flags
// DO-ONCE's codec does not already say, the confidence gate and the attack
// force (shown only when written; without it the force is 1).
func linkFields(p toulmin.EdgePacket, labels map[uint64]string) []Field {
	f := []Field{{Key: "src", Refs: refs([]toulmin.TID{p.Src}, labels)}}
	if p.DstFunc != nil {
		f = append(f, funcField("dst", *p.DstFunc, labels))
	} else {
		f = append(f, Field{Key: "dst", Refs: refs([]toulmin.TID{p.DstTID}, labels)})
	}
	if fl := flagsText(p.Flags); fl != "" {
		f = append(f, Field{Key: "flags", Text: fl})
	}
	if p.Gate != nil {
		f = append(f, Field{Key: "gate", Text: fmt.Sprintf("%s %d%%", gateOps[p.Gate.Op&3], p.Gate.Percent)})
	}
	if p.Force != nil {
		f = append(f, Field{Key: "force", Text: strconv.FormatFloat(p.Force.Float(), 'f', -1, 64)})
	}
	return f
}
