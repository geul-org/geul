//ff:func feature=delta type=decoder control=iteration dimension=1
//ff:what NewLocalIDs — the ID assignment state from a previous snapshot and its NEXT-LOCAL-ID: HEAD and BODY Edge contents to their IDs, REFERENCE-NODE objects to their LOCALs, described nodes to theirs
package delta

import (
	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/inspect"
)

// NewLocalIDs reads prev, the previous version's snapshot (nil for a first
// version), and next, the NEXT-LOCAL-ID of its version list (0 when there is
// none). Each Edge of the HEAD (a SOURCE on the whole document and its
// QUALIFIERs) and of the BODY is remembered by its content (edgeKey: codec and
// everything after the Edge Id); a REFERENCE-NODE whose Subject is a LOCAL is
// also remembered by its Object, defining that LOCAL, and a LOCAL that is the
// Subject of TRIPLEs by what they say of it (describedNodes). The next new ID is
// above both next and every LOCAL prev uses, and at least 1.
func NewLocalIDs(prev []byte, next uint64) (*LocalIDs, error) {
	l := &LocalIDs{next: max(next, 1), edges: map[string][]uint64{}, defs: map[string][][2]uint64{}, nodes: map[string][]uint64{}, used: map[uint64]bool{}}
	if prev == nil {
		return l, nil
	}
	ws, err := core.ReadWords(prev)
	if err != nil {
		return nil, err
	}
	doc, err := core.SplitDocument(ws)
	if err != nil {
		return nil, err
	}
	d, _ := inspect.Decode(prev)
	l.next = max(l.next, inspect.MaxLocalID(d)+1)
	for _, r := range append(append([]core.Record{}, doc.Head...), doc.Body...) {
		l.remember(r)
	}
	l.nodes = describedNodes(doc.Body)
	return l, nil
}
