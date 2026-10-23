//ff:func feature=codecspec type=decoder control=sequence
//ff:what TestCollect — a declared codec with named fields and enums; ISSUER targets, missing FIELDs and repeats are skipped
package codecspec

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestCollect(t *testing.T) {
	body := append(sampleBody(),
		core.EncodeDeclare(core.Declare{EdgeID: 11, Target: core.TargetWord(core.KindEdge, core.EdgeTombstone)}),
		core.EncodeDeclare(core.Declare{EdgeID: 12, Target: core.TargetWord(core.KindValue, core.CodecIssuer1Word), Issuer: []core.Word{1}, Format: 2}),
		core.EncodeDeclare(core.Declare{EdgeID: 13, Target: core.TargetWord(core.KindEdge, core.EdgeDeprecation), Fields: []uint64{2, 99}}),
	)
	s := Collect(body)
	fs := s[core.TargetWord(core.KindEdge, core.EdgeTombstone)]
	if len(s) != 1 || len(fs) != 3 || fs[0].Name != "doc" || fs[2].Width != 1 || fs[2].Enums["0x0001"] != "license" {
		t.Fatalf("got %+v", s)
	}
}
