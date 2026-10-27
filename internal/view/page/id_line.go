//ff:func feature=page type=helper control=sequence
//ff:what idLine — the dotted line from a LOCAL's ellipse to the hub of the Edge it names
package page

import (
	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/inspect"
)

// idLine returns the line joining ellipse end to the hub of the Edge that
// record field p names, when p is a LOCAL naming one of hubs: that ellipse is
// the Edge's Edge Id. ok is false for any other field.
func idLine(p inspect.Part, end string, hubs map[uint64]bool) (GraphEdge, bool) {
	r, _, _ := core.DecodeRecord(p.Words)
	id, err := core.DecodeLocalRef(r)
	hub := anchorOf(id)
	return GraphEdge{ID: end + "-id", Source: end, Target: hub, Kind: "ID", Class: "id", Row: hub}, err == nil && hubs[id]
}
