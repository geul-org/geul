//ff:func feature=geulmd type=encoder control=sequence
//ff:what TestRenderHead — ID-WIDTH declared after IMPORT-DOC still governs it
package geulmd

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestRenderHead(t *testing.T) {
	head := []core.Record{
		{Kind: core.KindMeta, Codec: core.MetaImportDoc, Payload: []core.Word{0, 3, 1, 0, 0, 0, 2}},
		{Kind: core.KindMeta, Codec: core.MetaIDWidth, Payload: []core.Word{0x0400}},
	}
	lines, ctx := renderHead(head)
	if len(lines) != 2 || lines[0] != "IMPORT-DOC #3 0x0001 0x0000000000000002" || ctx.importWidth != 2 || ctx.idWidth != 1 {
		t.Fatalf("got %q %+v", lines, ctx)
	}
}
