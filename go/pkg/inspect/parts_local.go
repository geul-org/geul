//ff:func feature=inspect type=decoder control=iteration dimension=1
//ff:what partsLocal — the largest LOCAL ID among the whole-record parts of a layout
package inspect

import "github.com/geul-org/geul/go/pkg/core"

// partsLocal returns the largest LOCAL Reference ID among ps's record parts
// (0 when none).
func partsLocal(ps []Part) uint64 {
	var max uint64
	for _, p := range ps {
		if !p.Record {
			continue
		}
		r, _, err := core.DecodeRecord(p.Words)
		if err != nil {
			continue
		}
		if id, err := core.DecodeLocalRef(r); err == nil {
			max = maxOf(max, id)
		}
	}
	return max
}
