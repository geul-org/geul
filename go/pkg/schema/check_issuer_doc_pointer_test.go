//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what TestCheckIssuerDocPointer — Issuer widths, document ID 0 and Quantifier fail R5; valid forms pass
package schema

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestCheckIssuerDocPointer(t *testing.T) {
	ref := func(c, q uint8, p ...core.Word) rec {
		return recOf(core.Record{Kind: core.KindReference, Codec: c<<3 | q, Payload: p})
	}
	iss, doc := core.RefIssuerPointer, core.RefDocPointer
	cases := []struct {
		r    rec
		c    uint8
		want string
	}{
		{ref(iss, 0, 0x40), iss, ""},
		{ref(iss, 0, 1, 2), iss, ""},
		{ref(iss, 0, 1, 2, 3), iss, "R5"},
		{ref(iss, 0), iss, "R5"},
		{ref(iss, 4, 0x40), iss, "R5"},
		{ref(doc, 0, 0, 0xFFFF, 0, 0, 1, 0, 1), doc, ""},
		{ref(doc, 0, 0, 0xFFFF, 0, 0, 1, 0, 0), doc, ""},
		{ref(doc, 0, 0, 0, 0, 0, 0, 0, 1), doc, "R5"},
		{ref(doc, 0, 0, 0, 1, 0, 0, 0, 1, 0, 1), doc, "R5"},
		{ref(doc, 1, 0, 0xFFFF, 0, 0, 1, 0, 1), doc, "R5"},
		{ref(doc, 1, 0, 0, 0, 0, 0, 0, 1), doc, "R5,R5"},
	}
	for i, tc := range cases {
		if got := rulesOf(checkIssuerDocPointer(tc.r, tc.c)); got != tc.want {
			t.Fatalf("%d: %q want %q", i, got, tc.want)
		}
	}
}
