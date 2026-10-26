//ff:func feature=delta type=encoder control=sequence
//ff:what TestLocalIDsEdge — an equal statement keeps its ID, duplicates in turn, a changed or new one gets the next ID
package delta

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestLocalIDsEdge(t *testing.T) {
	prev := encodeDoc(nil, []core.Record{bindOf(1, 9, "a"), bindOf(2, 9, "a"), bindOf(3, 9, "b")})
	l, _ := NewLocalIDs(prev, 0)
	subj, text := core.EncodeLocalRef(9), func(s string) core.Record { return core.EncodeBytesValue(core.CodecUTF8TextEven, []byte(s)) }
	if a, b := l.Edge(core.EdgeBind, words(subj, text("a"))), l.Edge(core.EdgeBind, words(subj, text("a"))); a != 1 || b != 2 {
		t.Fatalf("duplicates: %d %d", a, b)
	}
	if c := l.Edge(core.EdgeBind, words(subj, text("c"))); c != 10 || l.Next() != 11 {
		t.Fatalf("new: %d %d", c, l.Next())
	}
}
