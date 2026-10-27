//ff:func feature=page type=builder control=iteration dimension=1
//ff:what TestNodeFields — every flagged field in layout order; an empty node has none
package page

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/toulmin"
)

func TestNodeFields(t *testing.T) {
	q, v, c := toulmin.VFix(8000), toulmin.VFix(-10000), toulmin.TID(3)
	p := toulmin.NodePacket{
		Role: toulmin.RoleExcept, Level: toulmin.LevelReview, Func: &toulmin.FuncRef{Alias: 2, Verb: 1}, Qual: &q, Check: &c,
		Claim: &toulmin.Claim{Subject: 4, Kind: toulmin.ObjNum, ObjNum: 5}, Verdict: &v,
		Flags: toulmin.NodeFlagWith | toulmin.NodeFlagGrounds | toulmin.NodeFlagStrike, With: []toulmin.TID{6}, Grounds: []toulmin.TID{7},
	}
	want := []string{"role", "level", "func", "qualifier", "checking", "claim", "object", "verdict", "with", "grounds", "strike"}
	f := nodeFields(p, nil)
	if len(f) != len(want) || f[0].Text != "except" || f[3].Text != "0.8" || f[7].Text != "-1" {
		t.Fatalf("got %+v", f)
	}
	for i, k := range want {
		if f[i].Key != k {
			t.Fatalf("field %d: %q want %q", i, f[i].Key, k)
		}
	}
	if f := nodeFields(toulmin.NodePacket{}, nil); f != nil {
		t.Fatalf("empty: %+v", f)
	}
}
