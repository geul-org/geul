//ff:func feature=page type=builder control=iteration dimension=1
//ff:what edgeGraph — the lines of one Edge in RDF style: subject through its hub into the object, other fields into the hub
package page

import (
	"strconv"

	"github.com/geul-org/gogeul/pkg/inspect"
)

// edgeGraph draws Edge e around its hub (anchor `e<id>`). Its record fields
// after the Edge Id are its ends (graphEnd); the arrow runs from the hub to
// headEnd (class "head"), every other end runs into the hub (class "tail"),
// each line named by its field. Raw fields are left out. A LOCAL end naming an
// Edge also gets its idLine. It returns the end nodes and the lines.
func edgeGraph(e inspect.Entry, hubs map[uint64]bool, rows []Row, at map[string]int) ([]GraphNode, []GraphEdge) {
	var ns []GraphNode
	var es, ids []GraphEdge
	hub := anchorOf(e.ID)
	for k, p := range inspect.Layout(e) {
		if p.Head || !p.Record || p.Name == "ID" {
			continue
		}
		n := graphEnd(e, k, p, rows, at)
		ns = append(ns, n)
		es = append(es, GraphEdge{ID: hub + "-" + strconv.Itoa(k), Source: n.ID, Target: hub, Kind: e.Name, Class: "tail", Label: p.Name, Row: hub})
		if l, ok := idLine(p, n.ID, hubs); ok {
			ids = append(ids, l)
		}
	}
	if h := headEnd(ns); h >= 0 {
		es[h].Source, es[h].Target, es[h].Class = hub, ns[h].ID, "head"
	}
	return ns, append(es, ids...)
}
