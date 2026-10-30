//ff:func feature=toulmin type=helper control=iteration dimension=1
//ff:what TestEdgeRoundTrip — LINK records round-trip REF and FuncRef dsts, DO/DO-ONCE, UNDO once, gates and attack forces
package toulmin

import (
	"reflect"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

// TestEdgeRoundTrip checks DecodeEdge(EncodeEdge(p)) == p for an ATTACKS REF dst,
// a DO-ONCE FuncRef dst with a gate, a plain DO, an UNDO with F_ONCE kept on the
// wire, a LABEL (TEXT dst), a GEUL-format relation at a 2-word SPAN width and
// ATTACKS with an attack force of 0.5, an explicit 1 and 0.
func TestEdgeRoundTrip(t *testing.T) {
	g := Gate{Op: GateAtLeast, Percent: 75}
	half, one, zero := VFix(5000), VFix(10000), VFix(0)
	cases := []EdgePacket{
		{ID: 10, Prop: PropAttacks, Src: 2, DstTID: 3, Span: Span{Off: 5, Len: 6}},
		{ID: 11, Prop: PropDo, Src: 2, DstFunc: &FuncRef{Alias: 4, Verb: 1}, Flags: EdgeFlagOnce | EdgeFlagGate, Gate: &g},
		{ID: 12, Prop: PropDo, Src: 2, DstFunc: &FuncRef{}},
		{ID: 13, Prop: PropUndo, Src: 2, DstFunc: &FuncRef{Alias: 4}, Flags: EdgeFlagOnce | EdgeFlagOrdered},
		{ID: 14, Prop: PropLabel, Src: 2, DstTID: 15},
		{ID: 16, Prop: PropRequires, Src: 1, DstTID: 2, Flags: EdgeFlagStrike},
		{ID: 17, Prop: PropAttacks, Src: 2, DstTID: 3, Flags: EdgeFlagForce, Force: &half, Span: Span{Off: 1, Len: 2}},
		{ID: 18, Prop: PropAttacks, Src: 2, DstTID: 3, Flags: EdgeFlagForce | EdgeFlagGate, Gate: &g, Force: &one},
		{ID: 19, Prop: PropAttacks, Src: 2, DstTID: 3, Flags: EdgeFlagForce, Force: &zero},
	}
	wires := []Wire{testWire, {Span: 2}}
	for k := 0; k < len(wires)*len(cases); k++ {
		w, i := wires[k/len(cases)], k%len(cases)
		in := cases[i]
		ws, err := EncodeEdge(in, w)
		if err != nil {
			t.Fatalf("%d: EncodeEdge: %v", i, err)
		}
		r, _, err := core.DecodeRecord(ws)
		if err != nil {
			t.Fatal(err)
		}
		got, err := DecodeEdge(r, w)
		if err != nil || !reflect.DeepEqual(got, in) {
			t.Fatalf("%d: got %+v %v, want %+v", i, got, err, in)
		}
	}
}
