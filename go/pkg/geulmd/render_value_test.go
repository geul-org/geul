//ff:func feature=geulmd type=encoder control=iteration dimension=1
//ff:what TestRenderValue — every Value codec renders as its literal
package geulmd

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
	"github.com/geul-org/geul/go/pkg/inspect"
)

func TestRenderValue(t *testing.T) {
	cases := map[string]inspect.Value{
		`"a\n"`:      {Codec: core.CodecUTF8TextEven | 1, Text: "a\n"},
		`lang"ko"`:   {Codec: core.CodecLanguageBCP47Even, Text: "ko"},
		"uint(7)":    {Codec: core.CodecUnsignedInteger, Uint: 7},
		"dec(8e-1)":  {Codec: core.CodecSignedDecimal, Coef: 8, Exp: -1},
		"udec(75e0)": {Codec: core.CodecUnsignedDecimal, Coef: 75},
	}
	for want, v := range cases {
		if got := renderValue(v); got != want {
			t.Fatalf("%s: got %q", want, got)
		}
	}
}
