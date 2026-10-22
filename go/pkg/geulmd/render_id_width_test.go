//ff:func feature=geulmd type=encoder control=sequence
//ff:what TestRenderIDWidth — widths render and land in the context; reserved bits fail
package geulmd

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestRenderIDWidth(t *testing.T) {
	ctx := headCtx{}
	if got, ok := renderIDWidth(core.Record{Payload: []core.Word{0x4C00}}, &ctx); !ok || got != "ID-WIDTH edge=2 var=1 import=8" || ctx.idWidth != 2 || ctx.importWidth != 8 {
		t.Fatalf("got %q %v %+v", got, ok, ctx)
	}
	if _, ok := renderIDWidth(core.Record{Payload: []core.Word{1}}, &ctx); ok {
		t.Fatal("reserved bits: want not ok")
	}
}
