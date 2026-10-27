//ff:func feature=inspect type=decoder control=iteration dimension=1
//ff:what TestControlLayout — PROSE ends in a SPAN, SECTION in a code word, CASE-* in a REF
package inspect

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/toulmin"
)

func TestControlLayout(t *testing.T) {
	s := toulmin.Span{Off: 1, Len: 2}
	cases := map[string]toulmin.ControlPacket{
		"span.len": {ID: 1, Code: toulmin.ControlProse, Span: &s},
		"section":  {ID: 2, Code: toulmin.ControlSection, Arg: 5},
		"case":     {ID: 3, Code: toulmin.ControlCaseBegin, Arg: 1},
	}
	for last, p := range cases {
		ws, _ := toulmin.EncodeControl(p, testWire)
		r, _, _ := core.DecodeRecord(ws)
		w := &walker{ws: r.Payload, sw: 1}
		controlLayout(w, p)
		if w.bad || w.pos != len(w.ws) || w.out[len(w.out)-1].Name != last {
			t.Fatalf("%s: %+v", last, w)
		}
	}
}
