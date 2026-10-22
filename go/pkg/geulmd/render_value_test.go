//ff:func feature=geulmd type=encoder control=iteration dimension=1
//ff:what TestRenderValue — every known Value codec renders; unknown codecs and bad payloads do not
package geulmd

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestRenderValue(t *testing.T) {
	sd, _ := core.EncodeDecimalValue(true, -1, 8)
	ud, _ := core.EncodeDecimalValue(false, 0, 75)
	cases := map[string]core.Record{
		`"a\n"`:      core.EncodeBytesValue(core.CodecUTF8TextEven, []byte("a\n")),
		`lang"ko"`:   core.EncodeBytesValue(core.CodecLanguageBCP47Even, []byte("ko")),
		"uint(7)":    core.EncodeUintValue(7),
		"dec(8e-1)":  sd,
		"udec(75e0)": ud,
	}
	for want, r := range cases {
		if got, ok := renderValue(r); !ok || got != want {
			t.Fatalf("%s: got %q %v", want, got, ok)
		}
	}
	if _, ok := renderValue(core.Record{Kind: core.KindValue, Codec: 0x60}); ok {
		t.Fatal("unknown codec: want not ok")
	}
	if _, ok := renderValue(core.Record{Kind: core.KindValue, Codec: 0x41}); ok {
		t.Fatal("bad odd text: want not ok")
	}
}
