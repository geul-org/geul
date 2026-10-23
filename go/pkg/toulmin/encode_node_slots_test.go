//ff:func feature=toulmin type=encoder control=iteration dimension=1
//ff:what TestEncodeNodeSlots — all fields encode in flag order; each flag without its field and a bad claim are errors
package toulmin

import "testing"

func TestEncodeNodeSlots(t *testing.T) {
	q, v, chk := VFix(8000), VFix(10000), TID(9)
	all := NodeFlagFunc | NodeFlagQual | NodeFlagCheck | NodeFlagClaim | NodeFlagVerdict | NodeFlagWith | NodeFlagGrounds
	p := NodePacket{Flags: all, Func: &FuncRef{Alias: 3}, Qual: &q, Check: &chk, Claim: &Claim{}, Verdict: &v, With: []TID{5}, Grounds: []TID{6}}
	got, err := encodeNodeSlots(nil, p, testWire)
	if err != nil || len(got) != 27 {
		t.Fatalf("got %d words %v", len(got), err)
	}
	for _, f := range []NodeFlags{NodeFlagFunc, NodeFlagQual, NodeFlagCheck, NodeFlagClaim, NodeFlagVerdict} {
		if _, err := encodeNodeSlots(nil, NodePacket{Flags: f}, testWire); err == nil {
			t.Fatalf("flag %d without field: want error", f)
		}
	}
	if _, err := encodeNodeSlots(nil, NodePacket{Flags: NodeFlagClaim, Claim: &Claim{Kind: 3}}, testWire); err == nil {
		t.Fatal("bad claim object kind: want error")
	}
}
