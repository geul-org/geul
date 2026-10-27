//ff:func feature=delta type=decoder control=iteration dimension=1
//ff:what deltaHead — a diff's HEAD split into the snapshot HEAD and the removed IDs; an error without exactly one valid DELTA-BASE
package delta

import (
	"errors"

	"github.com/geul-org/gogeul/pkg/core"
)

// deltaHead takes DELTA-BASE and DELTA-REMOVE out of a diff's HEAD and
// returns the rest (the snapshot's HEAD) and the removed IDs (addRemoved). A
// diff has exactly one DELTA-BASE and at most one DELTA-REMOVE, both well
// formed.
func deltaHead(rs []core.Record) ([]core.Record, map[uint64]bool, error) {
	var head []core.Record
	removed := map[uint64]bool{}
	bases, removes := 0, 0
	for _, r := range rs {
		var err error
		switch {
		case r.Kind == core.KindMeta && r.Codec == core.MetaDeltaBase:
			_, err = core.DecodeDeltaBase(r)
			bases++
		case r.Kind == core.KindMeta && r.Codec == core.MetaDeltaRemove:
			err = addRemoved(removed, r)
			removes++
		default:
			head = append(head, r)
		}
		if err != nil {
			return nil, nil, err
		}
	}
	if bases != 1 || removes > 1 {
		return nil, nil, errors.New("geul/delta: a diff needs one DELTA-BASE and at most one DELTA-REMOVE")
	}
	return head, removed, nil
}
