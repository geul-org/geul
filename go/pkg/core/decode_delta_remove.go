//ff:func feature=core type=decoder control=iteration dimension=1
//ff:what DecodeDeltaRemove — the LOCAL IDs of a DELTA-REMOVE Meta; an error for another record, no ID, a non-LOCAL record or a repeated ID
package core

import (
	"errors"
	"fmt"
)

// DecodeDeltaRemove reads a DELTA-REMOVE Meta: one or more LOCAL Reference
// records, each ID once.
func DecodeDeltaRemove(r Record) ([]uint64, error) {
	if r.Kind != KindMeta || r.Codec != MetaDeltaRemove {
		return nil, fmt.Errorf("geul/core: expected Meta DELTA-REMOVE, got kind %d codec 0x%02X", r.Kind, r.Codec)
	}
	var ids []uint64
	seen := map[uint64]bool{}
	for ws := r.Payload; len(ws) > 0; {
		rec, n, err := DecodeRecord(ws)
		if err != nil {
			return nil, err
		}
		id, err := DecodeLocalRef(rec)
		if err != nil || seen[id] {
			return nil, fmt.Errorf("geul/core: DELTA-REMOVE needs distinct LOCAL References (%v)", err)
		}
		seen[id] = true
		ids = append(ids, id)
		ws = ws[n:]
	}
	if len(ids) == 0 {
		return nil, errors.New("geul/core: DELTA-REMOVE removes nothing")
	}
	return ids, nil
}
