//ff:func feature=toulmin type=encoder control=iteration dimension=1
//ff:what TestEncodeEdge — DO-ONCE golden (F_ONCE off the wire) plus every EncodeEdge error
package toulmin

import (
	"reflect"
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

// TestEncodeEdge pins a `do bank.withdraw once` LINK (codec 0x72, flag word 0)
// and checks the errors: provenance, unknown relation, ID/src/dst overflow, a
// FuncRef relation without DstFunc, a gate flag without a Gate, a bad gate and a
// SPAN overflow.
func TestEncodeEdge(t *testing.T) {
	got, err := EncodeEdge(EdgePacket{ID: 5, Prop: PropDo, Src: 2, DstFunc: &FuncRef{Alias: 3, Verb: 0xC200}, Flags: EdgeFlagOnce}, testWire)
	if want := []core.Word{0x13F2, 5, 2, 3, 0xC200, 0, 0, 0}; err != nil || !reflect.DeepEqual(got, want) {
		t.Fatalf("do once: %04X %v want %04X", got, err, want)
	}
	bad := []EdgePacket{
		{Prop: PropProvenance},
		{Prop: 0x0FFF},
		{Prop: PropAttacks, ID: 0x10000},
		{Prop: PropAttacks, Src: 0x10000},
		{Prop: PropAttacks, DstTID: 0x10000},
		{Prop: PropDo},
		{Prop: PropDo, DstFunc: &FuncRef{Alias: 0x10000}},
		{Prop: PropAttacks, Flags: EdgeFlagGate},
		{Prop: PropAttacks, Flags: EdgeFlagGate, Gate: &Gate{Percent: 101}},
		{Prop: PropAttacks, Span: Span{Len: 0x10000}},
	}
	for i, p := range bad {
		if _, err := EncodeEdge(p, testWire); err == nil {
			t.Fatalf("bad %d: want error", i)
		}
	}
}
