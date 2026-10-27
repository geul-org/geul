//ff:func feature=schema type=rule control=sequence
//ff:what TestHeadSources — the Edge Ids of the HEAD's document SOURCEs, not their QUALIFIERs or Meta
package schema

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestHeadSources(t *testing.T) {
	head := []rec{
		recOf(core.Record{Kind: core.KindMeta, Codec: core.MetaSpecVersion, Payload: []core.Word{1}}),
		recOf(edgeOf(core.EdgeSource, 3, 0x0088, 0, 0x00A0, 42)),
		recOf(edgeOf(core.EdgeQualifier, 4, 0x00A8, 813, 0x0088, 3, 0x00A0, 42)),
		recOf(edgeOf(core.EdgeSource, 7, 0x0088, 0, core.Absent)),
	}
	got := headSources(head)
	if len(got) != 2 || !got[3] || !got[7] {
		t.Fatalf("%v", got)
	}
	if headSources(nil) != nil {
		t.Fatal("no HEAD, no sources")
	}
}
