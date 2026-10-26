//ff:func feature=delta type=decoder control=sequence
//ff:what TestAddRemoved — the IDs are marked; a malformed DELTA-REMOVE errors and marks nothing
package delta

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestAddRemoved(t *testing.T) {
	m := map[uint64]bool{}
	if err := addRemoved(m, core.EncodeDeltaRemove([]uint64{3, 4})); err != nil || !m[3] || !m[4] {
		t.Fatalf("%v %v", m, err)
	}
	if err := addRemoved(map[uint64]bool{}, core.EncodeDeltaRemove(nil)); err == nil {
		t.Fatal("empty: want error")
	}
}
