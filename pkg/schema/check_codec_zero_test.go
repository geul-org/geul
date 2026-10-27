//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what TestCheckCodecZero — named slot markers as records are Errors; reserved markers and Codec 0000 with payload are Warns; the low bits are never a Quantifier
package schema

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestCheckCodecZero(t *testing.T) {
	ref := func(codec uint8, p ...core.Word) rec {
		return recOf(core.Record{Kind: core.KindReference, Codec: codec, Payload: p})
	}
	cases := []struct {
		r    rec
		sev  Severity
		part string
	}{
		{ref(0), Error, "자리 표지"},
		{ref(1), Error, "자리 표지"},
		{ref(2), Error, "자리 표지"},
		{ref(5), Warn, "자리 표지"},
		{ref(5, 1), Warn, "Codec"},
	}
	for i, c := range cases {
		fs := checkReference(c.r, nil)
		if len(fs) != 1 || fs[0].Rule != "R1" || fs[0].Severity != c.sev || fs[0].Where.Part != c.part {
			t.Fatalf("case %d: %+v", i, fs)
		}
	}
}
