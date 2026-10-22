//ff:func feature=geulmd type=encoder control=sequence
//ff:what TestNodeLine — every flag-gated field appears in flag order
package geulmd

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/toulmin"
)

func TestNodeLine(t *testing.T) {
	q, v, chk := toulmin.VFix(8000), toulmin.VFix(-10000), toulmin.TID(4)
	p := toulmin.NodePacket{TID: 2, Role: 1, Level: 1, Func: &toulmin.FuncRef{Alias: 3}, Qual: &q, Check: &chk,
		Claim: &toulmin.Claim{}, Verdict: &v, With: []toulmin.TID{5}, Grounds: nil,
		Flags: toulmin.NodeFlagWith | toulmin.NodeFlagGrounds | toulmin.NodeFlagStrike, Span: toulmin.Span{Len: 1}}
	want := "COUNTER-DECLARE #2 role=1 level=1 func=#3/0x0000 qual=0.8 check=#4 claim=#0/0x0000/#0/0x0000 verdict=-1 with=#5 grounds=# strike @0+1"
	if got := nodeLine("COUNTER-DECLARE", p); got != want {
		t.Fatalf("got  %q\nwant %q", got, want)
	}
}
