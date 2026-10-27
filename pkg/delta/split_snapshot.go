//ff:func feature=delta type=decoder control=iteration dimension=1
//ff:what splitSnapshot — a document as HEAD and Edge-Id-keyed BODY; an error when a BODY record is not an Edge with a LOCAL Edge Id
package delta

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// splitSnapshot splits b for diffing. Only a document whose BODY records are
// all Edges with a LOCAL Edge Id can be diffed (meta/lineage/DELTA-BASE).
func splitSnapshot(b []byte) (snapshot, error) {
	ws, err := core.ReadWords(b)
	if err != nil {
		return snapshot{}, err
	}
	doc, err := core.SplitDocument(ws)
	if err != nil {
		return snapshot{}, err
	}
	s := snapshot{head: doc.Head, body: doc.Body}
	for i, r := range doc.Body {
		id, _, ok := edgeID(r)
		if !ok {
			return snapshot{}, fmt.Errorf("geul/delta: BODY record %d is not an Edge with a LOCAL Edge Id", i+1)
		}
		s.ids = append(s.ids, id)
	}
	return s, nil
}
