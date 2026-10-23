//ff:func feature=inspect type=decoder control=iteration dimension=1
//ff:what edgeIDs — give every Edge entry without an Edge Id its leading LOCAL one
package inspect

import "github.com/geul-org/geul/go/pkg/core"

// edgeIDs sets the Edge Id of each Edge entry whose typed content gave none:
// every Edge's payload opens with a LOCAL Reference whose ID is its Edge Id
// (GEUL t002 OVERVIEW Edge). An Edge without one keeps 0.
func edgeIDs(entries []Entry) {
	for i := range entries {
		e := &entries[i]
		if e.Record.Kind != core.KindEdge || e.ID != 0 {
			continue
		}
		if r, _, err := core.DecodeRecord(e.Record.Payload); err == nil {
			e.ID, _ = core.DecodeLocalRef(r)
		}
	}
}
