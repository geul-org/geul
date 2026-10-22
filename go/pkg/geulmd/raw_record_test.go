//ff:func feature=geulmd type=encoder control=sequence
//ff:what TestRawRecord — direct codec, ISSUER codec and empty payload forms
package geulmd

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestRawRecord(t *testing.T) {
	if got := rawRecord(core.Record{Kind: core.KindEdge, Codec: 0x61, Payload: []core.Word{1}}); got != "RECORD EDGE 0x61 : 0x0001" {
		t.Fatalf("direct: %q", got)
	}
	if got := rawRecord(core.Record{Kind: core.KindEdge, Codec: 0x7C, Issuer: []core.Word{0}, Format: 0xC001}); got != "RECORD EDGE 0x7c 0x0000 0xc001 :" {
		t.Fatalf("issuer: %q", got)
	}
}
