//ff:func feature=page type=builder control=sequence
//ff:what TestBackrefMap — each referenced Edge Id lists its referrers in document order
package page

import (
	"slices"
	"testing"
)

func TestBackrefMap(t *testing.T) {
	back := backrefMap(fixtureDoc(t).Body)
	if !slices.Equal(back[1], []uint64{12, 2, 13, 14, 15}) || !slices.Equal(back[2], []uint64{14}) || !slices.Equal(back[3], []uint64{12}) {
		t.Fatalf("got %v", back)
	}
}
