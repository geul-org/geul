//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what TestDeltaBaseOf — a diff's base version; a snapshot, odd bytes and a broken document are not diffs
package resolve

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/delta"
)

func TestDeltaBaseOf(t *testing.T) {
	diff, _ := delta.Diff(edgeDoc(1, "a"), edgeDoc(1, "a", "b"), 3)
	if v, ok := deltaBaseOf(diff); !ok || v != 3 {
		t.Fatalf("%d %v", v, ok)
	}
	for _, b := range [][]byte{edgeDoc(1, "a"), {1}, {0, 5}} {
		if _, ok := deltaBaseOf(b); ok {
			t.Fatal("not a diff: want false")
		}
	}
}
