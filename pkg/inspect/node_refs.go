//ff:func feature=inspect type=helper control=sequence
//ff:what nodeRefs — a NODE's references: function ALIAS, checked case, claim subject/object, with and grounds TEXTs
package inspect

import "github.com/geul-org/gogeul/pkg/toulmin"

// nodeRefs returns the REF fields of a NODE in layout order: the FuncRef ALIAS,
// the F_CHECK case, the Claim subject and REF object, then the F_WITH and
// F_GROUNDS TEXT references.
func nodeRefs(p toulmin.NodePacket) []uint64 {
	var ids []toulmin.TID
	if p.Func != nil {
		ids = append(ids, p.Func.Alias)
	}
	if p.Check != nil {
		ids = append(ids, *p.Check)
	}
	if p.Claim != nil {
		ids = append(ids, p.Claim.Subject)
	}
	if p.Claim != nil && p.Claim.Kind == toulmin.ObjTID {
		ids = append(ids, p.Claim.ObjTID)
	}
	ids = append(ids, p.With...)
	return appendRefs(nil, append(ids, p.Grounds...)...)
}
