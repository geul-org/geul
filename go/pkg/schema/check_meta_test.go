//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what TestCheckMeta — each Meta branch and the reserved code
package schema

import (
	"github.com/geul-org/geul/go/pkg/core"
	"testing"
)

func TestCheckMeta(t *testing.T) {
	raw := func(c uint8, p ...core.Word) rec {
		return recOf(core.Record{Kind: core.KindMeta, Codec: c, Payload: p})
	}
	cases := []struct {
		r    rec
		want string
	}{
		{raw(core.MetaDocID, 1), "M2"},
		{raw(core.MetaImportDoc2Word, 1), "M3"},
		{raw(core.MetaSpecVersion), "M4"},
		{raw(core.MetaProfile), "M4"},
		{raw(core.MetaDefaultLanguage), "M4"},
		{raw(core.MetaEffectiveFrom), "M5"},
		{raw(metaParent), "M6"},
		{raw(0x15), "M1"},
		{raw(0x20), ""},
	}
	for i, c := range cases {
		if got := rulesOf(checkMeta(c.r)); got != c.want {
			t.Fatalf("case %d: %q", i, got)
		}
	}
}
