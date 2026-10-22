//ff:func feature=geulmd type=encoder control=sequence
//ff:what TestRenderBody — structured lines when lossless, raw when the typed form would change the words
package geulmd

import (
	"strings"
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
	"github.com/geul-org/geul/go/pkg/toulmin"
)

func TestRenderBody(t *testing.T) {
	ctx := headCtx{idWidth: 1, tangeul: true}
	node, _ := toulmin.EncodeNode(toulmin.NodePacket{TID: 1, Type: toulmin.NodeCase}, toulmin.Wire{ID: 1, Span: 1})
	r, _, _ := core.DecodeRecord(node)
	// A non-trimmed qualifier decimal (80e-2) decodes fine but re-encodes as 8e-1.
	untrimmed := core.Record{Kind: core.KindValue, Codec: core.CodecSignedDecimal, Payload: []core.Word{0xFFFF, 0xFFFE, 80}}
	lines := renderBody([]core.Record{r, untrimmed}, ctx)
	if lines[0] != "CASE-DECLARE #1 role=0 level=0 @0+0" || lines[1] != "VALUE dec(80e-2)" {
		t.Fatalf("got %q", lines)
	}
	raw := core.Record{Kind: core.KindEdge, Codec: toulmin.CodecCaseDeclare, Payload: []core.Word{1, 0, 0}}
	if got := renderBody([]core.Record{raw}, ctx); !strings.HasPrefix(got[0], "RECORD EDGE 0x60") {
		t.Fatalf("malformed node must be raw: %q", got)
	}
}
