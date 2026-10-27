//ff:func feature=page type=helper control=sequence
//ff:what TestIDLine — a LOCAL naming an Edge joins its ellipse to that hub; other LOCALs and Values none
package page

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/inspect"
)

func TestIDLine(t *testing.T) {
	parts := inspect.Layout(bindEntry(2, 1, "a"))
	l, ok := idLine(parts[2], "n0", map[uint64]bool{1: true})
	if !ok || l != (GraphEdge{ID: "n0-id", Source: "n0", Target: "e1", Kind: "ID", Class: "id", Row: "e1"}) {
		t.Fatalf("got %+v %v", l, ok)
	}
	if _, ok := idLine(parts[2], "n0", nil); ok {
		t.Fatal("not an Edge")
	}
	if _, ok := idLine(parts[3], "x", map[uint64]bool{1: true}); ok {
		t.Fatal("value")
	}
}
