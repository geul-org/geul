//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what TestCheckReference — reserved Quantifier and Codec, pointer, issuer and declaration forms
package schema

import (
	"github.com/geul-org/gogeul/pkg/core"
	"testing"
)

func TestCheckReference(t *testing.T) {
	ref := func(codec uint8, p ...core.Word) rec {
		return recOf(core.Record{Kind: core.KindReference, Codec: codec, Payload: p})
	}
	cases := map[string]rec{
		"R1":    ref(0x15, 1),
		"R1,R1": ref(0x05, 1),
		"R3":    ref(0x40, 9, 1),
		"R4":    ref(0x60, 1),
		"":      ref(0x10, 1),
		"R2":    ref(0x18, 0, 1),
		"R5":    ref(0x39, 0, 0, 0, 0, 0, 0, 1),
	}
	for want, r := range cases {
		if got := rulesOf(checkReference(r, nil)); got != want {
			t.Fatalf("codec 0x%02X: %q want %q", r.R.Codec, got, want)
		}
	}
}
