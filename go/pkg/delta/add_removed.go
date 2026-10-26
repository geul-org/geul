//ff:func feature=delta type=decoder control=iteration dimension=1
//ff:what addRemoved — the IDs of a DELTA-REMOVE Meta added to a set; its decoding error
package delta

import "github.com/geul-org/geul/go/pkg/core"

// addRemoved decodes DELTA-REMOVE r and marks each of its IDs in removed.
func addRemoved(removed map[uint64]bool, r core.Record) error {
	ids, err := core.DecodeDeltaRemove(r)
	for _, id := range ids {
		removed[id] = true
	}
	return err
}
