//ff:func feature=delta type=helper control=sequence
//ff:what TestUnused — previous Edge Ids no new statement took, in previous order
package delta

import (
	"slices"
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestUnused(t *testing.T) {
	prev := encodeDoc(nil, []core.Record{bindOf(1, 9, "a"), bindOf(2, 9, "b"), bindOf(3, 9, "c")})
	l, _ := NewLocalIDs(prev, 0)
	l.Edge(core.EdgeBind, words(core.EncodeLocalRef(9), core.EncodeBytesValue(core.CodecUTF8TextEven, []byte("b"))))
	if got := l.Unused(); !slices.Equal(got, []uint64{1, 3}) {
		t.Fatalf("got %v", got)
	}
}
