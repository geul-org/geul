//ff:func feature=inspect type=helper control=sequence
//ff:what TestAppendRefs — zero ids are skipped, order kept
package inspect

import (
	"slices"
	"testing"
)

func TestAppendRefs(t *testing.T) {
	if got := appendRefs([]uint64{7}, 0, 3, 0, 1); !slices.Equal(got, []uint64{7, 3, 1}) {
		t.Fatalf("got %v", got)
	}
}
