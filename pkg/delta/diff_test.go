//ff:func feature=delta type=encoder control=iteration dimension=1
//ff:what TestDiff — a diff holds only changed and new Edges plus DELTA-BASE/REMOVE and applies back to next; reordering and bad inputs fail
package delta

import (
	"bytes"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestDiff(t *testing.T) {
	head := []core.Record{core.EncodeSpecVersion(2)}
	base := encodeDoc(head, []core.Record{bindOf(1, 9, "a"), bindOf(2, 9, "b"), bindOf(3, 9, "c")})
	next := encodeDoc(head, []core.Record{bindOf(2, 9, "B"), bindOf(3, 9, "c"), bindOf(4, 9, "d")})
	d, err := Diff(base, next, 1)
	if err != nil {
		t.Fatal(err)
	}
	x, _ := splitSnapshot(d)
	if len(x.body) != 2 || x.ids[0] != 2 || x.ids[1] != 4 || len(x.head) != 3 {
		t.Fatalf("diff body %v head %d", x.ids, len(x.head))
	}
	if got, _ := Apply(base, d); !bytes.Equal(got, next) {
		t.Fatal("round trip")
	}
	same, _ := Diff(base, base, 1)
	if s, _ := splitSnapshot(same); len(s.body) != 0 || len(s.head) != 2 {
		t.Fatalf("no change: %d %d", len(s.body), len(s.head))
	}
	reordered := encodeDoc(head, []core.Record{bindOf(3, 9, "c"), bindOf(1, 9, "a"), bindOf(2, 9, "b")})
	for name, c := range map[string][2][]byte{"reordered": {base, reordered}, "base": {[]byte{1}, next}, "next": {base, []byte{1}}} {
		if _, err := Diff(c[0], c[1], 1); err == nil {
			t.Fatalf("%s: want error", name)
		}
	}
}
