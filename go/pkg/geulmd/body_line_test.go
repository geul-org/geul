//ff:func feature=geulmd type=encoder control=iteration dimension=1
//ff:what TestBodyLine — every interpreted BODY content renders its line; bare records are raw
package geulmd

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
	"github.com/geul-org/geul/go/pkg/inspect"
	"github.com/geul-org/geul/go/pkg/toulmin"
)

func TestBodyLine(t *testing.T) {
	src := inspect.Source{TextPacket: toulmin.TextPacket{TID: 9, Bytes: []byte("# t\n")}, SpanWidth: 1}
	cases := map[string]inspect.Entry{
		"VALUE uint(5)":                       {Data: inspect.Value{Codec: core.CodecUnsignedInteger, Uint: 5}},
		`SOURCE #9 span=1 "# t\n"`:            {Data: src},
		`TEXT #10 "n"`:                        {Data: toulmin.TextPacket{TID: 10, Bytes: []byte("n")}},
		"SECTION #11 5":                       {Name: "SECTION", Data: toulmin.ControlPacket{ID: 11, Code: toulmin.ControlSection, Arg: 5}},
		"CASE-DECLARE #1 role=0 level=0 @0+4": {Name: "CASE-DECLARE", Data: toulmin.NodePacket{TID: 1, Span: toulmin.Span{Len: 4}}},
		"LABEL #12 #1 -> #10 @0+4":            {Name: "LABEL", Data: toulmin.EdgePacket{ID: 12, Src: 1, DstTID: 10, Span: toulmin.Span{Len: 4}}},
		"RECORD EDGE 0x61 : 0x000d":           {Name: "PREDICATE-DECLARE", Record: core.Record{Kind: core.KindEdge, Codec: 0x61, Payload: []core.Word{13}}},
	}
	for want, e := range cases {
		if got := bodyLine(e); got != want {
			t.Fatalf("got %q want %q", got, want)
		}
	}
}
