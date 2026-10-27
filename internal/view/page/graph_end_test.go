//ff:func feature=page type=helper control=sequence
//ff:what TestGraphEnd — a Value as a literal, any other record as its Node row's ellipse (name when unlabelled)
package page

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/inspect"
)

func TestGraphEnd(t *testing.T) {
	e := bindEntry(1, 7, "a")
	parts := inspect.Layout(e)
	rows := []Row{{Anchor: "n0", Name: "LOCAL"}}
	at := map[string]int{}
	if n := graphEnd(e, 3, parts[3], rows, at); n.ID != "e1-3v" || n.Class != "lit" || n.Label != `"a"` || n.Row != "e1" {
		t.Fatalf("value: %+v", n)
	}
	if n := graphEnd(e, 2, parts[2], rows, at); n != (GraphNode{ID: "n0", Label: "LOCAL", Kind: "LOCAL", Class: "res"}) {
		t.Fatalf("ref: %+v", n)
	}
	rows[0].Label = "x"
	if n := graphEnd(e, 2, parts[2], rows, at); n.Label != "x" {
		t.Fatalf("labelled: %+v", n)
	}
}
