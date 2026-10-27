//ff:func feature=page type=builder control=sequence
//ff:what TestEdgeGraph — subject only: hub → it; subject and object: subject → hub → object; a LOCAL naming an Edge adds its ID line
package page

import (
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/inspect"
)

func TestEdgeGraph(t *testing.T) {
	e := bindEntry(1, 7, "a")
	e.Name = "BIND"
	rows, at := refNodes([]inspect.Entry{e}, nil)
	ns, es := edgeGraph(e, map[uint64]bool{1: true}, rows, at)
	if len(ns) != 2 || ns[0].ID != "n0" || ns[1].ID != "e1-3v" {
		t.Fatalf("nodes: %+v", ns)
	}
	if len(es) != 2 || es[0] != (GraphEdge{ID: "e1-2", Source: "e1", Target: "n0", Kind: "BIND", Class: "head", Label: "subject", Row: "e1"}) ||
		es[1] != (GraphEdge{ID: "e1-3", Source: "e1-3v", Target: "e1", Kind: "BIND", Class: "tail", Label: "name", Row: "e1"}) {
		t.Fatalf("edges: %+v", es)
	}
	enc := func(r core.Record) []core.Word { ws, _ := core.EncodeRecord(r); return ws }
	p := slices.Concat(enc(core.EncodeLocalRef(5)), enc(core.EncodeLocalRef(7)), enc(core.EncodeLocalRef(1)))
	rel := inspectEntry(core.Record{Kind: core.KindEdge, Codec: core.EdgeBind, Payload: p}, nil)
	rel.ID = 5
	rows, at = refNodes([]inspect.Entry{rel}, nil)
	_, es = edgeGraph(rel, map[uint64]bool{1: true, 5: true}, rows, at)
	if len(es) != 3 || es[0].Source != "n0" || es[0].Class != "tail" || es[1].Source != "e5" || es[1].Target != "n1" || es[1].Class != "head" || es[2].ID != "n1-id" {
		t.Fatalf("binary: %+v", es)
	}
}
