//ff:func feature=toulmin type=helper control=iteration dimension=1
//ff:what TestNodeRoundTrip — NODE records round-trip every slot combination at 1- and 2-word widths
package toulmin

import (
	"reflect"
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

// TestNodeRoundTrip encodes NODE records with every flag-gated field (FuncRef,
// qualifier, checking REF, each Claim object kind, verdict, with list) under
// rule and non-rule codecs and checks DecodeNode restores them exactly, at the
// narrow widths and at 2-word Edge Id / SPAN widths.
func TestNodeRoundTrip(t *testing.T) {
	q, v := VFix(8000), VFix(-2500)
	chk := TID(9)
	all := NodeFlagFunc | NodeFlagQual | NodeFlagCheck | NodeFlagClaim | NodeFlagVerdict | NodeFlagWith | NodeFlagGrounds
	cases := []NodePacket{
		{TID: 1, Type: NodeCase, Span: Span{Off: 3, Len: 4}},
		{TID: 2, Type: NodeRule, Role: RoleCounter, Level: LevelFail, Flags: all,
			Func: &FuncRef{Alias: 3, Verb: 0x7400}, Qual: &q, Check: &chk, Verdict: &v, With: []TID{5, 6}, Grounds: []TID{10},
			Claim: &Claim{Subject: 4, Predicate: 7, Kind: ObjTID, ObjTID: 8, Mode: 1}},
		{TID: 3, Type: NodeTerm, Role: 2, Flags: NodeFlagClaim | NodeFlagStrike,
			Claim: &Claim{Kind: ObjNum, ObjNum: 70000}},
		{TID: 4, Type: NodeField, Flags: NodeFlagClaim, Claim: &Claim{Kind: ObjSpan, ObjSpan: Span{Off: 1, Len: 2}}},
		{TID: 5, Type: NodeRule, Role: RoleExcept, Flags: NodeFlagWith, With: []TID{}},
	}
	wires := []Wire{testWire, {ID: 2, Span: 2}}
	for k := 0; k < len(wires)*len(cases); k++ {
		w, i := wires[k/len(cases)], k%len(cases)
		in := cases[i]
		ws, err := EncodeNode(in, w)
		if err != nil {
			t.Fatalf("%d: EncodeNode: %v", i, err)
		}
		r, n, err := core.DecodeRecord(ws)
		if err != nil || n != len(ws) {
			t.Fatalf("%d: not one record: %v", i, err)
		}
		got, err := DecodeNode(r, w)
		if err != nil || !reflect.DeepEqual(got, in) {
			t.Fatalf("%d (wire %+v): got %+v %v, want %+v", i, w, got, err, in)
		}
	}
}
