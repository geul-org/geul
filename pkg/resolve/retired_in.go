//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what retiredIn — the first LOCAL ID new in a version that is below NEXT-LOCAL-ID: dropped before or skipped
package resolve

import "github.com/geul-org/gogeul/pkg/inspect"

// retiredIn returns the first LOCAL ID of b (inspect.LocalIDs, document order)
// that prev, the previous stored version, does not use and that is below
// next, the version list's NEXT-LOCAL-ID. A new statement takes its ID from
// NEXT-LOCAL-ID and a dropped ID is never used again (meta/lineage/
// NEXT-LOCAL-ID), so such an ID was dropped before or skipped. prev nil means
// no stored version, so every ID below next counts. #0 is the document itself
// (the Target of a HEAD SOURCE, edge/SOURCE), never an issued ID, so it is
// never retired.
func retiredIn(prev, b []byte, next uint64) (uint64, bool) {
	pd, _ := inspect.Decode(prev)
	bd, _ := inspect.Decode(b)
	kept := map[uint64]bool{}
	for _, id := range inspect.LocalIDs(pd) {
		kept[id] = true
	}
	for _, id := range inspect.LocalIDs(bd) {
		if id != 0 && id < next && !kept[id] {
			return id, true
		}
	}
	return 0, false
}
