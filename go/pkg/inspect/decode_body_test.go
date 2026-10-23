//ff:func feature=inspect type=decoder control=sequence
//ff:what TestDecodeBody — a SOURCE's SPAN width reaches the records after it
package inspect

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
	"github.com/geul-org/geul/go/pkg/toulmin"
)

func TestDecodeBody(t *testing.T) {
	w := toulmin.Wire{Span: 2}
	src, _ := toulmin.EncodeSource(toulmin.TextPacket{TID: 1}, w)
	node, _ := toulmin.EncodeNode(toulmin.NodePacket{TID: 2, Type: toulmin.NodeCase, Span: toulmin.Span{Off: 70000}}, w)
	rs, _, _ := core.DecodeRecord(src)
	rn, _, _ := core.DecodeRecord(node)
	out := decodeBody([]core.Record{rs, rn}, true)
	if p, ok := out[1].Data.(toulmin.NodePacket); len(out) != 2 || !ok || p.Span.Off != 70000 {
		t.Fatalf("got %+v", out)
	}
}
