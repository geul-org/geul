//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what TestCheckReference — reserved Quantifier and Codec, pointer, issuer and declaration forms
package schema

import (
	"github.com/geul-org/geul/go/pkg/core"
	"testing"
)

func TestCheckReference(t *testing.T) {
	ref := func(codec uint8, p ...core.Word) rec {
		return recOf(core.Record{Kind: core.KindReference, Codec: codec, Payload: p})
	}
	cases := map[string]rec{
		"R1":    ref(0x45, 1),
		"R1,R1": ref(0x65, 1),
		"R3":    ref(0x00, 9, 1),
		"R4":    ref(0x20, 1),
		"":      ref(0x40, 1),
		"R2":    ref(0x48, 0, 1),
	}
	for want, r := range cases {
		if got := rulesOf(checkReference(r, nil)); got != want {
			t.Fatalf("codec 0x%02X: %q want %q", r.R.Codec, got, want)
		}
	}
}
