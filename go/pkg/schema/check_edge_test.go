//ff:func feature=schema type=rule control=sequence
//ff:what TestCheckEdge — missing Edge Id, reserved codec, bad and good version list layouts
package schema

import (
	"github.com/geul-org/geul/go/pkg/core"
	"testing"
)

func TestCheckEdge(t *testing.T) {
	if fs := checkEdge(recOf(core.Record{Kind: core.KindEdge, Codec: 0x03})); rulesOf(fs) != "E1,E3" {
		t.Fatalf("bad: %+v", fs)
	}
	if fs := checkEdge(recOf(edgeOf(core.EdgeDeprecation, 1, 0, 0, 0, 5, 0))); rulesOf(fs) != "E4" {
		t.Fatalf("layout: %+v", fs)
	}
	if fs := checkEdge(recOf(edgeOf(core.EdgeDeprecation, 1, 0, 0, 0, 5, 0, 1))); fs != nil {
		t.Fatalf("ok: %+v", fs)
	}
}
