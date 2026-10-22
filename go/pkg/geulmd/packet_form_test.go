//ff:func feature=geulmd type=encoder control=iteration dimension=1
//ff:what TestPacketForm — TEXT, control, NODE and LINK forms; malformed controls and unsupported records fail
package geulmd

import (
	"slices"
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
	"github.com/geul-org/geul/go/pkg/toulmin"
)

func TestPacketForm(t *testing.T) {
	w := toulmin.Wire{ID: 1, Span: 1}
	txt, _ := toulmin.EncodeText(toulmin.TextPacket{TID: 1, Bytes: []byte("a")}, w)
	ctl, _ := toulmin.EncodeControl(toulmin.ControlPacket{ID: 2, Code: toulmin.ControlSection, Arg: 3}, w)
	node, _ := toulmin.EncodeNode(toulmin.NodePacket{TID: 3, Type: toulmin.NodeCase}, w)
	link, _ := toulmin.EncodeEdge(toulmin.EdgePacket{ID: 4, Prop: toulmin.PropAttacks, Src: 3, DstTID: 3}, w)
	for _, ws := range [][]core.Word{txt, ctl, node, link} {
		r, _, _ := core.DecodeRecord(ws)
		ref, _ := toulmin.RecordRef(r)
		_, enc, err := packetForm(toulminName(ref), ref, r, w)
		if err != nil || !slices.Equal(enc, ws) {
			t.Fatalf("%04X: %v", ws, err)
		}
	}
	badCtl := core.Record{Kind: core.KindEdge, Codec: 0x7C, Issuer: []core.Word{0}, Format: toulmin.FormatSection}
	if _, _, err := packetForm("SECTION", toulmin.CodecRef{Codec: 0x7C, Format: toulmin.FormatSection}, badCtl, w); err == nil {
		t.Fatal("malformed control: want error")
	}
	if _, _, err := packetForm("PREDICATE-DECLARE", toulmin.CodecRef{Codec: toulmin.CodecPredicateDeclare}, core.Record{}, w); err == nil {
		t.Fatal("unsupported record: want error")
	}
}
