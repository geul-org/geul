//ff:func feature=delta type=decoder control=sequence
//ff:what TestAddedEdges — only Edge Ids missing from before, in order
package delta

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestAddedEdges(t *testing.T) {
	n, _ := splitSnapshot(encodeDoc(nil, []core.Record{bindOf(1, 9, "a"), bindOf(2, 9, "b"), bindOf(3, 9, "c")}))
	cs := addedEdges(n, map[uint64]bool{2: true})
	if len(cs) != 2 || cs[0].ID != 1 || cs[1].ID != 3 || cs[0].Op != '+' {
		t.Fatalf("%+v", cs)
	}
}
