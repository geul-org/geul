//ff:func feature=toulmin type=decoder control=iteration dimension=1
//ff:what TestRecordRef — direct and GEUL-format Edge records resolve; other kinds and Issuers fail
package toulmin

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestRecordRef(t *testing.T) {
	if r, err := RecordRef(core.Record{Kind: core.KindEdge, Codec: CodecRun}); err != nil || r.Codec != CodecRun || r.Format != 0 {
		t.Fatalf("direct: %+v %v", r, err)
	}
	if r, err := RecordRef(core.Record{Kind: core.KindEdge, Codec: 0x7C, Issuer: []core.Word{0}, Format: FormatText}); err != nil || r.Format != FormatText {
		t.Fatalf("format: %+v %v", r, err)
	}
	bad := []core.Record{
		{Kind: core.KindValue},
		{Kind: core.KindEdge, Codec: 0x7C, Issuer: []core.Word{1}},
		{Kind: core.KindEdge, Codec: 0x7D, Issuer: []core.Word{0, 0}},
	}
	for i, b := range bad {
		if _, err := RecordRef(b); err == nil {
			t.Fatalf("bad %d: want error", i)
		}
	}
}
