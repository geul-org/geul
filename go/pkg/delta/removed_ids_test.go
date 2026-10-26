//ff:func feature=delta type=helper control=sequence
//ff:what TestRemovedIDs — base IDs next lacks, in base order
package delta

import (
	"slices"
	"testing"
)

func TestRemovedIDs(t *testing.T) {
	if got := removedIDs([]uint64{1, 2, 3, 4}, []uint64{4, 2, 9}); !slices.Equal(got, []uint64{1, 3}) {
		t.Fatalf("got %v", got)
	}
}
