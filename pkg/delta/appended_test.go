//ff:func feature=delta type=decoder control=sequence
//ff:what TestAppended — the left-over diff Edges in diff order
package delta

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestAppended(t *testing.T) {
	x := snapshot{body: []core.Record{bindOf(4, 9, "d"), bindOf(5, 9, "e"), bindOf(6, 9, "f")}, ids: []uint64{4, 5, 6}}
	got := appended(x, map[uint64]core.Record{6: {}, 4: {}})
	if len(got) != 2 || got[0].Payload[1] != 4 || got[1].Payload[1] != 6 {
		t.Fatalf("%+v", got)
	}
}
