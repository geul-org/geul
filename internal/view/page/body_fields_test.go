//ff:func feature=page type=builder control=iteration dimension=1
//ff:what TestBodyFields — every decoded content type yields its first field; declared records show spec fields, bare records raw fields
package page

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/codecspec"
	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/inspect"
	"github.com/geul-org/gogeul/pkg/toulmin"
)

func TestBodyFields(t *testing.T) {
	cases := map[string]inspect.Entry{
		"value":   {Data: inspect.Value{Codec: core.CodecUnsignedInteger}},
		"n":       inspectEntry(core.Record{Kind: core.KindValue, Codec: 0x30, Payload: []core.Word{1}}, []codecspec.Field{{Name: "n", Width: 1}}),
		"bytes":   {Data: inspect.Source{TextPacket: toulmin.TextPacket{Bytes: []byte("ab")}, SpanWidth: 1}},
		"text":    {Data: toulmin.TextPacket{Bytes: []byte("t")}},
		"section": {Data: toulmin.ControlPacket{Code: toulmin.ControlSection, Arg: 1}},
		"role":    {Data: toulmin.NodePacket{Role: toulmin.RoleCounter}},
		"src":     {Data: toulmin.EdgePacket{Src: 1}},
		"issuer":  {Record: core.Record{Issuer: []core.Word{0}}},
	}
	for want, e := range cases {
		if f := bodyFields(e, nil); len(f) == 0 || f[0].Key != want {
			t.Fatalf("%s: %+v", want, f)
		}
	}
}
