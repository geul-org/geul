//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what TestCheckValue — each Value branch and the reserved code
package schema

import (
	"github.com/geul-org/gogeul/pkg/core"
	"testing"
)

func TestCheckValue(t *testing.T) {
	v := func(c uint8, p ...core.Word) rec {
		return recOf(core.Record{Kind: core.KindValue, Codec: c, Payload: p})
	}
	cases := []struct {
		r    rec
		want string
	}{
		{v(0x7C), ""},
		{v(0x00, 2), "V2"},
		{v(0x01, 0, 1), "V3"},
		{v(0x02, 0, 1), "V3"},
		{v(0x03, 1, 2, 3), "V2"},
		{v(0x11, 0), "V3"},
		{v(0x21, 1), "V2"},
		{v(0x60, 1), "V2"},
		{v(0x62, 1), "V2"},
		{v(0x62, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0), ""},
		{v(0x63), "V1"},
		{v(0x41), "V4"},
		{v(0x25), "V4"},
		{v(0x24, 0x3139, 0x3734), ""},
		{v(0x27), "V4"},
		{v(0x30), "V1"},
		{v(0x04, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0), ""},
		{v(0x05), "V8"},
		{v(0x06), "V9"},
		{v(0x42, 0x4142), "V1"},
		{v(0x4A, 0x08C6, 0x656E), ""},
	}
	for i, c := range cases {
		if got := rulesOf(checkValue(c.r)); got != c.want {
			t.Fatalf("case %d: %q", i, got)
		}
	}
}
