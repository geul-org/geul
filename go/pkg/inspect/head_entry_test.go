//ff:func feature=inspect type=decoder control=iteration dimension=1
//ff:what TestHeadEntry — each known Meta gets typed Data; unknown and malformed Meta stay bare
package inspect

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestHeadEntry(t *testing.T) {
	lw, _ := core.EncodeRecord(core.EncodeBytesValue(core.CodecLanguageBCP47Even, []byte("en")))
	typed := []core.Record{
		{Kind: core.KindMeta, Codec: core.MetaDefaultLanguage, Payload: lw},
		{Kind: core.KindMeta, Codec: core.MetaDocName, Payload: []core.Word{1, 0, 0, 0, 2}},
		{Kind: core.KindMeta, Codec: core.MetaImportDoc2Word, Payload: append([]core.Word{4, 1, 1, 0, 0, 0, 2}, make([]core.Word, 16)...)},
	}
	for _, r := range typed {
		if e := headEntry(r); e.Data == nil || e.Name == "" {
			t.Fatalf("0x%02x: %+v", r.Codec, e)
		}
	}
	if e := headEntry(core.Record{Kind: core.KindMeta, Codec: core.MetaSpecVersion}); e.Data != nil || e.Name != "SPEC-VERSION" {
		t.Fatalf("malformed: %+v", e)
	}
	if e := headEntry(core.Record{Kind: core.KindMeta, Codec: 0x20, Payload: []core.Word{1}}); e.Data != nil || e.Name != "" {
		t.Fatalf("unknown: %+v", e)
	}
}
