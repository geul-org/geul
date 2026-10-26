//ff:func feature=resolve type=resolver control=sequence
//ff:what TestListParent — a PARENT digest is read; no PARENT and a malformed one are false
package resolve

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestListParent(t *testing.T) {
	var p [32]byte
	p[0], p[31] = 0xAB, 0xCD
	b := encodeVersionList(nil, &p, 1, 0)
	ws, _ := core.ReadWords(b)
	d, _ := core.SplitDocument(ws)
	if h, ok := listParent(d.Head); !ok || h != p {
		t.Fatalf("got %x %v", h, ok)
	}
	if _, ok := listParent(nil); ok {
		t.Fatal("none: want false")
	}
	bad := core.Record{Kind: core.KindMeta, Codec: 0x30, Payload: []core.Word{0x0080}}
	if _, ok := listParent([]core.Record{bad}); ok {
		t.Fatal("malformed: want false")
	}
}
