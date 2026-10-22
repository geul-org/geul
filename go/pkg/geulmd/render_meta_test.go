//ff:func feature=geulmd type=encoder control=iteration dimension=1
//ff:what TestRenderMeta — each interpreted Meta form and the raw fallback
package geulmd

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
	"github.com/geul-org/geul/go/pkg/inspect"
)

func TestRenderMeta(t *testing.T) {
	cases := map[string]inspect.Entry{
		`PROFILE "other-1.0"`:                     {Name: "PROFILE", Data: inspect.Value{Codec: core.CodecUTF8TextEven, Text: "other-1.0"}},
		"ID-WIDTH edge=2 var=1 import=4":          {Data: inspect.Widths{Edge: 2, Var: 1, Import: 4}},
		"DOC-NAME 0x0001 0x0000000000000002":      {Data: inspect.DocName{Issuer: []core.Word{1}, Doc: []core.Word{0, 0, 0, 2}}},
		"IMPORT-DOC #4 0x0001 0x0000000000000002": {Data: inspect.ImportDoc{ID: 4, Issuer: []core.Word{1}, Doc: []core.Word{0, 0, 0, 2}}},
		"RECORD META 0x20 : 0x0001":               {Record: core.Record{Kind: core.KindMeta, Codec: 0x20, Payload: []core.Word{1}}},
	}
	for want, e := range cases {
		if got := renderMeta(e); got != want {
			t.Fatalf("got %q want %q", got, want)
		}
	}
}
