//ff:func feature=schema type=rule control=sequence
//ff:what TestCheckIssuerRef — short payload fails R4; a full one checks its target
package schema

import (
	"github.com/geul-org/geul/go/pkg/core"
	"testing"
)

func TestCheckIssuerRef(t *testing.T) {
	ref := func(p ...core.Word) rec {
		return recOf(core.Record{Kind: core.KindReference, Codec: 4 << 3, Payload: p})
	}
	if fs := checkIssuerRef(ref(1, 0, 0, 0, 2), 4); rulesOf(fs) != "R4" {
		t.Fatalf("short: %+v", fs)
	}
	if fs := checkIssuerRef(ref(1, 0, 0, 0, 2, 7), 4); fs != nil {
		t.Fatalf("ok: %+v", fs)
	}
}
