//ff:func feature=inspect type=decoder control=iteration dimension=1
//ff:what TestParseValue — every known codec decodes (byte strings incl. URI, EDTF, GEDCOM-DATE, language text; coordinates; integer, decimal, time, digest); unknown codecs and bad payloads do not
package inspect

import (
	"strings"
	"testing"
	"time"

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
		{core.EncodeLanguageText(core.EncodeBytesValue(core.CodecLanguageBCP47Even, []byte("ko")), []byte("우주")), Value{Codec: core.CodecUTF8LanguageTextEven, Text: "우주", Lang: "ko"}},
		{core.EncodeBytesValue(0x44, []byte("http://x")), Value{Codec: 0x44, Text: "http://x"}},
		{core.EncodeBytesValue(core.CodecEDTFEven, []byte("1974")), Value{Codec: core.CodecEDTFEven, Text: "1974"}},
		{core.EncodeBytesValue(core.CodecGEDCOMDateEven, []byte("JULIAN 1300")), Value{Codec: core.CodecGEDCOMDateEven | 1, Text: "JULIAN 1300"}},
		{core.EncodeUintValue(7), Value{Codec: core.CodecUnsignedInteger, Uint: 7}},
		{core.EncodeCoordinate(core.Coordinate{Lat: 52.5, Lon: 13.4, U: 44}), Value{Codec: core.CodecCoordinateEarth, Lat: 52.5, Lon: 13.4, U: 44}},
		{sd, Value{Codec: core.CodecSignedDecimal, Exp: -1, Coef: 8}},
		{ud, Value{Codec: core.CodecUnsignedDecimal, Coef: 75}},
		{core.EncodePosixSeconds(86400), Value{Codec: core.CodecPosixSeconds, Time: time.Unix(86400, 0).UTC()}},
		{core.Record{Kind: core.KindValue, Codec: core.CodecGregorianDate, Payload: []core.Word{0, 1}}, Value{Codec: core.CodecGregorianDate, Time: time.Unix(86400, 0).UTC()}},
		{core.Record{Kind: core.KindValue, Codec: core.CodecSHA256Digest, Payload: make([]core.Word, 16)}, Value{Codec: core.CodecSHA256Digest, Text: "0x" + strings.Repeat("0", 64)}},
	}
	for _, c := range cases {
		if got, ok := ParseValue(c.r); !ok || got != c.want {
			t.Fatalf("got %+v %v want %+v", got, ok, c.want)
		}
	}
	for _, c := range []uint8{0x60, 0x30, 0x42} {
		if _, ok := ParseValue(core.Record{Kind: core.KindValue, Codec: c}); ok {
			t.Fatalf("0x%02X: want not ok", c)
		}
	}
	if _, ok := ParseValue(core.Record{Kind: core.KindValue, Codec: core.CodecSHA256Digest, Payload: []core.Word{1}}); ok {
		t.Fatal("short digest: want not ok")
	}
	if _, ok := ParseValue(core.Record{Kind: core.KindValue, Codec: 0x41}); ok {
		t.Fatal("bad odd text: want not ok")
	}
}
