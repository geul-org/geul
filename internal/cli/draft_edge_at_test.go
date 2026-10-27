//ff:func feature=cli type=helper control=sequence
//ff:what TestDraftEdgeAt — an Edge Id found in the BODY, then the HEAD; -1 when in neither
package cli

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/notation"
)

func TestDraftEdgeAt(t *testing.T) {
	e := func(id uint64) core.Record { return edgeRecord(id, notation.Edge{Codec: core.EdgeSource}) }
	d := &draft{head: []core.Record{core.EncodeSpecVersion(core.SpecVersion), e(2)}, body: []core.Record{e(5), e(7)}}
	if part, i := d.edgeAt(7); part != &d.body || i != 1 {
		t.Fatalf("BODY: %d", i)
	}
	if part, i := d.edgeAt(2); part != &d.head || i != 1 {
		t.Fatalf("HEAD: %d", i)
	}
	if _, i := d.edgeAt(9); i != -1 {
		t.Fatalf("none: %d", i)
	}
}
