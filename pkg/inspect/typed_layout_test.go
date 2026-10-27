//ff:func feature=inspect type=decoder control=iteration dimension=1
//ff:what TestTypedLayout — each TOULMIN record content walks its whole payload; other content has no typed layout
package inspect

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/toulmin"
)

func TestTypedLayout(t *testing.T) {
	src, _ := toulmin.EncodeSource(toulmin.TextPacket{TID: 1, Bytes: []byte("a")}, testWire)
	txt, _ := toulmin.EncodeText(toulmin.TextPacket{TID: 2, Bytes: []byte("b")}, testWire)
	ctl, _ := toulmin.EncodeControl(toulmin.ControlPacket{ID: 3, Code: toulmin.ControlCaseEnd, Arg: 1}, testWire)
	node, _ := toulmin.EncodeNode(toulmin.NodePacket{TID: 4, Type: toulmin.NodeCase}, testWire)
	edge, _ := toulmin.EncodeEdge(toulmin.EdgePacket{ID: 5, Prop: toulmin.PropAttacks, Src: 1, DstTID: 4}, testWire)
	cases := []struct {
		ws   []core.Word
		data any
	}{
		{src, Source{TextPacket: toulmin.TextPacket{TID: 1}, SpanWidth: 1}},
		{txt, toulmin.TextPacket{TID: 2}},
		{ctl, toulmin.ControlPacket{ID: 3, Code: toulmin.ControlCaseEnd, Arg: 1}},
		{node, toulmin.NodePacket{TID: 4, Type: toulmin.NodeCase}},
		{edge, toulmin.EdgePacket{ID: 5, Prop: toulmin.PropAttacks, Src: 1, DstTID: 4}},
	}
	for i, c := range cases {
		r, _, _ := core.DecodeRecord(c.ws)
		w := &walker{ws: r.Payload, sw: 1}
		if !typedLayout(w, c.data) || w.bad || w.pos != len(w.ws) {
			t.Fatalf("case %d: %+v", i, w)
		}
	}
	if typedLayout(&walker{}, Value{}) {
		t.Fatal("Value: want no typed layout")
	}
}
