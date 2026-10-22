//ff:func feature=geulmd type=encoder control=iteration dimension=1
//ff:what TestRenderMeta — typed Meta, a non-tangeul profile, and raw fallback for unknown or malformed Meta
package geulmd

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestRenderMeta(t *testing.T) {
	vw, _ := core.EncodeRecord(core.EncodeBytesValue(core.CodecUTF8TextEven, []byte("other-1.0")))
	ctx := headCtx{importWidth: 1}
	if got := renderMeta(core.Record{Kind: core.KindMeta, Codec: core.MetaProfile, Payload: vw}, &ctx); got != `PROFILE "other-1.0"` || ctx.tangeul {
		t.Fatalf("profile: %q %+v", got, ctx)
	}
	if got := renderMeta(core.Record{Kind: core.KindMeta, Codec: core.MetaSpecVersion}, &ctx); got != "RECORD META 0x10 :" {
		t.Fatalf("malformed: %q", got)
	}
	if got := renderMeta(core.Record{Kind: core.KindMeta, Codec: 0x20, Payload: []core.Word{1}}, &ctx); got != "RECORD META 0x20 : 0x0001" {
		t.Fatalf("unknown: %q", got)
	}
	lw, _ := core.EncodeRecord(core.EncodeBytesValue(core.CodecLanguageBCP47Even, []byte("en")))
	typed := map[string]core.Record{
		`DEFAULT-LANGUAGE lang"en"`:               {Kind: core.KindMeta, Codec: core.MetaDefaultLanguage, Payload: lw},
		"ID-WIDTH edge=1 var=1 import=1":          {Kind: core.KindMeta, Codec: core.MetaIDWidth, Payload: []core.Word{0}},
		"DOC-NAME 0x0001 0x0000000000000002":      {Kind: core.KindMeta, Codec: core.MetaDocName, Payload: []core.Word{1, 0, 0, 0, 2}},
		"IMPORT-DOC #4 0x0001 0x0000000000000002": {Kind: core.KindMeta, Codec: core.MetaImportDoc, Payload: []core.Word{4, 1, 0, 0, 0, 2}},
	}
	for want, r := range typed {
		if got := renderMeta(r, &headCtx{importWidth: 1}); got != want {
			t.Fatalf("got %q want %q", got, want)
		}
	}
}
