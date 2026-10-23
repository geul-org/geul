//ff:func feature=toulmin type=decoder control=sequence
//ff:what TestDecodeNodeSlots — every flag-gated field decodes in order; an absent field is an error
package toulmin

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestDecodeNodeSlots(t *testing.T) {
	all := NodeFlagFunc | NodeFlagQual | NodeFlagCheck | NodeFlagClaim | NodeFlagVerdict | NodeFlagWith | NodeFlagGrounds
	ws := []core.Word{
		0x00C0, 3, 0x7400, // FuncRef
		0x0990, 0xFFFF, 0xFFFF, 8, // qualifier 0.8
		0x00C0, 9, // checking REF
		0x00C0, 4, 7, 0x00C0, 8, 1, // Claim
		0x0990, 0, 0, 1, // verdict 1
		0x0882, 1, 0x00C0, 5, // with [5]
		0x0882, 1, 0x00C0, 6, // grounds [6]
	}
	p, pos := NodePacket{Flags: all}, 0
	if err := decodeNodeSlots(ws, &pos, &p, ObjTID, testWire); err != nil || pos != len(ws) || *p.Check != 9 || p.Claim.Mode != 1 || len(p.With) != 1 || len(p.Grounds) != 1 {
		t.Fatalf("got %+v %v pos=%d", p, err, pos)
	}
	q := NodePacket{Flags: NodeFlagQual}
	pos = 0
	if err := decodeNodeSlots(nil, &pos, &q, ObjTID, testWire); err == nil {
		t.Fatal("missing qualifier: want error")
	}
}
