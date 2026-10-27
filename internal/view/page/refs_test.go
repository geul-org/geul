//ff:func feature=page type=helper control=sequence
//ff:what TestRefs — links carry labels; 0 (none) is left out
package page

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/toulmin"
)

func TestRefs(t *testing.T) {
	got := refs([]toulmin.TID{0, 4, 5}, map[uint64]string{4: "a"})
	if len(got) != 2 || got[0] != (Ref{ID: 4, Label: "a"}) || got[1] != (Ref{ID: 5}) {
		t.Fatalf("got %+v", got)
	}
}
