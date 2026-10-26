//ff:func feature=delta type=decoder control=iteration dimension=1
//ff:what TestDeltaHead — DELTA-BASE and DELTA-REMOVE are taken out; none, two bases, two removals and malformed ones fail
package delta

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestDeltaHead(t *testing.T) {
	sv := core.EncodeSpecVersion(2)
	head, removed, err := deltaHead([]core.Record{sv, core.EncodeDeltaBase(1), core.EncodeDeltaRemove([]uint64{3})})
	if err != nil || len(head) != 1 || !removed[3] {
		t.Fatalf("%v %v %v", head, removed, err)
	}
	for i, rs := range [][]core.Record{
		{sv},
		{core.EncodeDeltaBase(1), core.EncodeDeltaBase(2)},
		{core.EncodeDeltaBase(1), core.EncodeDeltaRemove([]uint64{1}), core.EncodeDeltaRemove([]uint64{2})},
		{core.EncodeDeltaBase(0)},
		{core.EncodeDeltaBase(1), core.EncodeDeltaRemove(nil)},
	} {
		if _, _, err := deltaHead(rs); err == nil {
			t.Fatalf("case %d: want error", i)
		}
	}
}
