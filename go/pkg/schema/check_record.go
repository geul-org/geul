//ff:func feature=schema type=rule control=selection
//ff:what checkRecord — the per-record rules of one HEAD or BODY record by Kind
package schema

import "github.com/geul-org/geul/go/pkg/core"

// checkRecord dispatches one record to its Kind's rules.
func checkRecord(r rec, imports map[uint64]bool) []Fact {
	switch r.R.Kind {
	case core.KindReference:
		return checkReference(r, imports)
	case core.KindValue:
		return checkValue(r)
	case core.KindEdge:
		return checkEdge(r)
	default:
		return checkMeta(r)
	}
}
