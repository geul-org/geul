//ff:func feature=delta type=decoder control=sequence
//ff:what TestBlame — each Edge of the last version with the version that added it and the one that last changed it; no versions and a non-snapshot fail
package delta

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestBlame(t *testing.T) {
	v1 := encodeDoc(nil, []core.Record{bindOf(1, 9, "a"), bindOf(2, 9, "b")})
	v2 := encodeDoc(nil, []core.Record{bindOf(1, 9, "a"), bindOf(2, 9, "B"), bindOf(3, 9, "c")})
	v3 := encodeDoc(nil, []core.Record{bindOf(1, 9, "a"), bindOf(3, 9, "c"), bindOf(4, 9, "d")})
	ls, err := Blame([]Version{{1, v1}, {2, v2}, {3, v3}})
	if err != nil || len(ls) != 3 {
		t.Fatalf("%+v %v", ls, err)
	}
	if ls[0].ID != 1 || ls[0].Added != 1 || ls[0].Changed != 1 || ls[1].ID != 3 || ls[1].Added != 2 || ls[2].Added != 3 {
		t.Fatalf("%+v", ls)
	}
	ls, _ = Blame([]Version{{1, v1}, {2, v2}})
	if ls[1].Added != 1 || ls[1].Changed != 2 {
		t.Fatalf("changed: %+v", ls[1])
	}
	if _, err := Blame(nil); err == nil {
		t.Fatal("none: want error")
	}
	if _, err := Blame([]Version{{1, []byte{1}}}); err == nil {
		t.Fatal("bad snapshot: want error")
	}
}
