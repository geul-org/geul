//ff:func feature=inspect type=decoder control=iteration dimension=1
//ff:what TestEdgeLayout — a gated FuncRef LINK, a plain REF LINK and an ATTACKS with force walk their whole payloads
package inspect

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/toulmin"
)

func TestEdgeLayout(t *testing.T) {
	g := toulmin.Gate{Op: toulmin.GateAtLeast, Percent: 75}
	half := toulmin.VFix(5000)
	cases := map[int]toulmin.EdgePacket{
		9: {ID: 1, Prop: toulmin.PropDo, Src: 2, DstFunc: &toulmin.FuncRef{Alias: 3}, Flags: toulmin.EdgeFlagGate, Gate: &g},
		6: {ID: 4, Prop: toulmin.PropAttacks, Src: 2, DstTID: 3},
		7: {ID: 5, Prop: toulmin.PropAttacks, Src: 2, DstTID: 3, Flags: toulmin.EdgeFlagForce, Force: &half},
	}
	for parts, p := range cases {
		ws, _ := toulmin.EncodeEdge(p, testWire)
		r, _, _ := core.DecodeRecord(ws)
		w := &walker{ws: r.Payload, sw: 1}
		edgeLayout(w, p)
		if w.bad || w.pos != len(w.ws) || len(w.out) != parts {
			t.Fatalf("%d: %+v", parts, w.out)
		}
	}
}
