//ff:func feature=delta type=decoder control=sequence
//ff:what TestBodyChanges — unchanged Edges give nothing; removed, changed and added are found by Edge Id
package delta

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestBodyChanges(t *testing.T) {
	s := func(rs ...core.Record) snapshot { x, _ := splitSnapshot(encodeDoc(nil, rs)); return x }
	if cs := bodyChanges(s(bindOf(1, 9, "a")), s(bindOf(1, 9, "a"))); len(cs) != 0 {
		t.Fatalf("same: %+v", cs)
	}
	cs := bodyChanges(s(bindOf(1, 9, "a"), bindOf(2, 9, "b")), s(bindOf(2, 9, "B"), bindOf(5, 9, "e")))
	if len(cs) != 3 || cs[0].Op != '-' || cs[0].ID != 1 || cs[1].Op != '~' || cs[2].Op != '+' || cs[2].ID != 5 {
		t.Fatalf("%+v", cs)
	}
}
