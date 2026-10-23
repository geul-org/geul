//ff:func feature=inspect type=decoder control=sequence
//ff:what TestNodeLayout — a NODE with every flag-gated field walks its whole payload in flag order
package inspect

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
	"github.com/geul-org/geul/go/pkg/toulmin"
)

func TestNodeLayout(t *testing.T) {
	q, v, chk := toulmin.VFix(8000), toulmin.VFix(10000), toulmin.TID(9)
	all := toulmin.NodeFlagFunc | toulmin.NodeFlagQual | toulmin.NodeFlagCheck | toulmin.NodeFlagClaim | toulmin.NodeFlagVerdict | toulmin.NodeFlagWith | toulmin.NodeFlagGrounds
	p := toulmin.NodePacket{TID: 1, Type: toulmin.NodeRule, Flags: all, Func: &toulmin.FuncRef{Alias: 3}, Qual: &q, Check: &chk, Claim: &toulmin.Claim{}, Verdict: &v, With: []toulmin.TID{5}, Grounds: []toulmin.TID{6, 7}}
	ws, _ := toulmin.EncodeNode(p, testWire)
	r, _, _ := core.DecodeRecord(ws)
	w := &walker{ws: r.Payload, sw: 1}
	nodeLayout(w, p)
	if w.bad || w.pos != len(w.ws) || len(w.out) != 18 || w.out[2].Name != "func.alias" {
		t.Fatalf("got %+v", w)
	}
}
