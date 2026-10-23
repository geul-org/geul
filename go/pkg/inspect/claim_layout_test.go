//ff:func feature=inspect type=decoder control=iteration dimension=1
//ff:what TestClaimLayout — a REF or Value object is one record, a SPAN object two words
package inspect

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
	"github.com/geul-org/geul/go/pkg/toulmin"
)

func TestClaimLayout(t *testing.T) {
	cases := map[int]toulmin.Claim{
		4: {Subject: 1, Kind: toulmin.ObjTID, ObjTID: 2},
		5: {Subject: 1, Kind: toulmin.ObjSpan, ObjSpan: toulmin.Span{Off: 1, Len: 2}},
	}
	for parts, c := range cases {
		p := toulmin.NodePacket{TID: 1, Type: toulmin.NodeRule, Flags: toulmin.NodeFlagClaim, Claim: &c}
		ws, _ := toulmin.EncodeNode(p, testWire)
		r, _, _ := core.DecodeRecord(ws)
		w := &walker{ws: r.Payload[3:], sw: 1}
		claimLayout(w, c)
		if w.bad || len(w.out) != parts {
			t.Fatalf("%+v: %+v", c, w)
		}
	}
}
