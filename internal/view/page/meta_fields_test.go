//ff:func feature=page type=builder control=sequence
//ff:what TestMetaFields — Value, DOC-ID, IMPORT-DOC latest and pinned, raw ISSUER-less Meta
package page

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/inspect"
)

func TestMetaFields(t *testing.T) {
	if f := metaFields(inspect.Entry{Data: inspect.Value{Codec: core.CodecUnsignedInteger, Uint: 3}}); len(f) != 1 || f[0].Text != "3" {
		t.Fatalf("value: %+v", f)
	}
	if f := metaFields(inspect.Entry{Data: inspect.DocID{Issuer: []core.Word{1}, Doc: []core.Word{0, 0, 0, 5}}}); len(f) != 2 || f[1].Text != "0x0000000000000005" {
		t.Fatalf("doc name: %+v", f)
	}
	imp := inspect.ImportDoc{ID: 2, Issuer: []core.Word{1}, Doc: []core.Word{0, 0, 0, 5}}
	if f := metaFields(inspect.Entry{Data: imp}); len(f) != 4 || f[0].Text != "#2" || f[3].Text != "latest" {
		t.Fatalf("import: %+v", f)
	}
	imp.Version = 1
	if f := metaFields(inspect.Entry{Data: imp}); len(f) != 4 || f[3].Text == "latest" {
		t.Fatalf("pinned: %+v", f)
	}
	if f := metaFields(inspect.Entry{Record: core.Record{Kind: core.KindMeta, Codec: 0x20}}); f != nil {
		t.Fatalf("raw: %+v", f)
	}
}
