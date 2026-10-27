//ff:func feature=page type=builder control=sequence
//ff:what nodeFields — every field a NODE's flags carry, in layout order
package page

import (
	"strconv"

	"github.com/geul-org/gogeul/pkg/toulmin"
)

// roleNames and levelNames name the NODE header codes (GEUL t002 TOULMIN NODE).
var (
	roleNames  = map[toulmin.RoleCode]string{1: "counter", 2: "except"}
	levelNames = map[toulmin.LevelCode]string{1: "fail", 2: "review"}
)

// nodeFields lays out a NODE: a non-zero role and level, then each field its
// flags carry — function, qualifier, checked case, claim, verdict, with,
// grounds — and the strike flag.
func nodeFields(p toulmin.NodePacket, labels map[uint64]string) []Field {
	var f []Field
	if p.Role != 0 {
		f = append(f, Field{Key: "role", Text: codeName(roleNames, p.Role)})
	}
	if p.Level != 0 {
		f = append(f, Field{Key: "level", Text: codeName(levelNames, p.Level)})
	}
	if p.Func != nil {
		f = append(f, funcField("func", *p.Func, labels))
	}
	if p.Qual != nil {
		f = append(f, Field{Key: "qualifier", Text: strconv.FormatFloat(p.Qual.Float(), 'f', -1, 64)})
	}
	if p.Check != nil {
		f = append(f, Field{Key: "checking", Refs: refs([]toulmin.TID{*p.Check}, labels)})
	}
	if p.Claim != nil {
		f = append(f, claimFields(*p.Claim, labels)...)
	}
	if p.Verdict != nil {
		f = append(f, Field{Key: "verdict", Text: strconv.FormatFloat(p.Verdict.Float(), 'f', -1, 64)})
	}
	if p.Flags&toulmin.NodeFlagWith != 0 {
		f = append(f, Field{Key: "with", Refs: refs(p.With, labels)})
	}
	if p.Flags&toulmin.NodeFlagGrounds != 0 {
		f = append(f, Field{Key: "grounds", Refs: refs(p.Grounds, labels)})
	}
	if p.Flags&toulmin.NodeFlagStrike != 0 {
		f = append(f, Field{Key: "strike"})
	}
	return f
}
