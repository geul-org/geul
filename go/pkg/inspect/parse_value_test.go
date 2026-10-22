//ff:func feature=inspect type=decoder control=iteration dimension=1
//ff:what TestParseValue — every known codec decodes; unknown codecs and bad payloads do not
package inspect

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestParseValue(t *testing.T) {
	sd, _ := core.EncodeDecimalValue(true, -1, 8)
	ud, _ := core.EncodeDecimalValue(false, 0, 75)
	cases := []struct {
		r    core.Record
		want Value
	}{
		{core.EncodeBytesValue(core.CodecUTF8TextEven, []byte("abc")), Value{Codec: core.CodecUTF8TextEven | 1, Text: "abc"}},
		{core.EncodeBytesValue(core.CodecLanguageBCP47Even, []byte("ko")), Value{Codec: core.CodecLanguageBCP47Even, Text: "ko"}},
		{core.EncodeUintValue(7), Value{Codec: core.CodecUnsignedInteger, Uint: 7}},
		{sd, Value{Codec: core.CodecSignedDecimal, Exp: -1, Coef: 8}},
		{ud, Value{Codec: core.CodecUnsignedDecimal, Coef: 75}},
	}
	for _, c := range cases {
		if got, ok := parseValue(c.r); !ok || got != c.want {
			t.Fatalf("got %+v %v want %+v", got, ok, c.want)
		}
	}
	if _, ok := parseValue(core.Record{Kind: core.KindValue, Codec: 0x60}); ok {
		t.Fatal("unknown codec: want not ok")
	}
	if _, ok := parseValue(core.Record{Kind: core.KindValue, Codec: 0x41}); ok {
		t.Fatal("bad odd text: want not ok")
	}
}
