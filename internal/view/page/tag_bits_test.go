//ff:func feature=page type=helper control=iteration dimension=1
//ff:what TestTagBits — Edge, Reference with Quantifier, a slot marker (no Quantifier), 1-word and extended Lengths, a fallback name and an unnamed codec
package page

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestTagBits(t *testing.T) {
	cases := map[core.Word]string{
		0x13FC: "000 Proposal\n10 EDGE\n0111 8 WORDs\n1111100 ISSUER-1WORD",
		0x0088: "000 Proposal\n00 REFERENCE\n0001 2 WORDs\n0001 LOCAL\n000 Identified",
		0x1801: "000 Proposal\n11 META\n0000 1 WORD\n0000001 DOC-2WORD",
		0x1761: "000 Proposal\n10 EDGE\n1110 Length 확장 1 WORD\n1100001 0x61",
		0x17E6: "000 Proposal\n10 EDGE\n1111 Length 확장 2 WORD\n1100110 0x66",
		0x0001: "000 Proposal\n00 REFERENCE\n0000 1 WORD\n0000 자리 표지\n001 표지 번호",
	}
	for w, want := range cases {
		if got := tagBits(w, ""); got != want {
			t.Fatalf("%04X: %q want %q", uint16(w), got, want)
		}
	}
	if got := tagBits(0x1463, "RULE-DECLARE"); got != "000 Proposal\n10 EDGE\n1000 9 WORDs\n1100011 RULE-DECLARE" {
		t.Fatalf("fallback: %q", got)
	}
}
