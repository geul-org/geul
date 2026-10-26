//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what listNext — the NEXT-LOCAL-ID a version list's HEAD holds; 0 when it has none
package resolve

import "github.com/geul-org/geul/go/pkg/core"

// listNext returns the next LOCAL ID a version list keeps in its HEAD
// (meta/lineage/NEXT-LOCAL-ID), or 0 when it has none or it does not decode.
func listNext(head []core.Record) uint64 {
	for _, r := range head {
		if n, err := core.DecodeNextLocalID(r); err == nil {
			return n
		}
	}
	return 0
}
