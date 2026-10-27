//ff:func feature=page type=builder control=sequence
//ff:what TestRefNodes — records inside Edges once each with their Edges as back-references and BIND names as labels; Edge Ids and Values left out
package page

import (
	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/inspect"
	"testing"
)

func TestRefNodes(t *testing.T) {
	body := []inspect.Entry{bindEntry(1, 7, "a"), bindEntry(2, 7, "b"), bindEntry(3, 1, "c"), inspectEntry(core.EncodeUintValue(5), nil)}
	rows, at := refNodes(body, map[uint64]string{2: "two"})
	if len(rows) != 2 || rows[0].Anchor != "n0" || rows[0].Label != "b" || len(rows[0].Back) != 2 || rows[0].Back[1].Label != "two" {
		t.Fatalf("got %+v", rows)
	}
	if rows[1].Label != "c" || len(rows[1].Back) != 1 || len(at) != 2 {
		t.Fatalf("edge ref: %+v", rows[1])
	}
}
