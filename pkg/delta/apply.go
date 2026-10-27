//ff:func feature=delta type=decoder control=iteration dimension=1
//ff:what Apply — the snapshot a diff makes from its base: HEAD from the diff, base BODY in order with removals and replacements, new Edges appended
package delta

import (
	"errors"
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// Apply rebuilds the snapshot that diff d makes from base (OVERVIEW 계보
// diff): the HEAD is d's HEAD without DELTA-BASE and DELTA-REMOVE; the BODY
// keeps base's order, drops the removed IDs, replaces an Edge in place when
// d has the same Edge Id and appends d's other Edges in d's order. d must
// carry DELTA-BASE (deltaHead); a removed ID that base lacks or that d also
// writes is an error. The caller checks the result's content hash against
// the version list.
func Apply(base, d []byte) ([]byte, error) {
	b, err := splitSnapshot(base)
	if err != nil {
		return nil, err
	}
	x, err := splitSnapshot(d)
	if err != nil {
		return nil, err
	}
	head, removed, err := deltaHead(x.head)
	if err != nil {
		return nil, err
	}
	writes := map[uint64]core.Record{}
	for i, id := range x.ids {
		if removed[id] {
			return nil, fmt.Errorf("geul/delta: Edge #%d is both removed and written", id)
		}
		writes[id] = x.body[i]
	}
	var body []core.Record
	for i, id := range b.ids {
		r, written := writes[id]
		switch {
		case removed[id]:
			delete(removed, id)
		case written:
			body = append(body, r)
			delete(writes, id)
		default:
			body = append(body, b.body[i])
		}
	}
	if len(removed) > 0 {
		return nil, errors.New("geul/delta: DELTA-REMOVE names an Edge Id the base does not have")
	}
	return encodeDoc(head, append(body, appended(x, writes)...)), nil
}
