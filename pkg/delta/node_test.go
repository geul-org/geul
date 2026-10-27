//ff:func feature=delta type=encoder control=sequence
//ff:what TestNode — a node described as before keeps its LOCAL, twins keep theirs in order, a taken or new description gets a new ID
package delta

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestNode(t *testing.T) {
	prev := encodeDoc(nil, []core.Record{
		tripleOf(1, 854, core.EncodeLocalRef(20), "u"),
		tripleOf(2, 854, core.EncodeLocalRef(21), "u"),
	})
	l, _ := NewLocalIDs(prev, 30)
	_, u, _ := tripleAbout(tripleOf(9, 854, core.EncodeLocalRef(1), "u"))
	about := [][]core.Word{u}
	if a, b, c := l.Node(about), l.Node(about), l.Node(about); a != 20 || b != 21 || c != 30 {
		t.Fatalf("%d %d %d", a, b, c)
	}
	l.used[23] = true
	l.nodes[nodeKey(nil)] = []uint64{23}
	if got := l.Node(nil); got != 31 {
		t.Fatalf("new: %d", got)
	}
}
