//ff:func feature=schema type=rule control=sequence
//ff:what TestCheckIssuerRef — short payload fails R4; a full one checks its target; document 0, target 0 and a non-canonical Issuer fail R4
package schema

import (
	"github.com/geul-org/geul/go/pkg/core"
	"testing"
)

func TestCheckIssuerRef(t *testing.T) {
	ref := func(p ...core.Word) rec {
		return recOf(core.Record{Kind: core.KindReference, Codec: core.RefExternalIssuer1Word << 3, Payload: p})
	}
	if fs := checkIssuerRef(ref(1, 0, 0, 0, 2), core.RefExternalIssuer1Word); rulesOf(fs) != "R4" {
		t.Fatalf("short: %+v", fs)
	}
	if fs := checkIssuerRef(ref(1, 0, 0, 0, 2, 7), core.RefExternalIssuer1Word); fs != nil {
		t.Fatalf("ok: %+v", fs)
	}
	if fs := checkIssuerRef(ref(1, 0, 0, 0, 2, 0), core.RefExternalIssuer1Word); rulesOf(fs) != "R4" {
		t.Fatalf("document itself: %+v", fs)
	}
	if fs := checkIssuerRef(ref(1, 0, 0, 0, 0, 7), core.RefExternalIssuer1Word); rulesOf(fs) != "R4" {
		t.Fatalf("document 0, target 7: %+v", fs)
	}
	if fs := checkIssuerRef(ref(1, 0, 0, 0, 0, 0), core.RefExternalIssuer1Word); rulesOf(fs) != "R4,R4" {
		t.Fatalf("issuer itself: %+v", fs)
	}
	if fs := checkIssuerRef(ref(1, 0, 0, 0, 2, 0, 7), core.RefExternalIssuer1Word); rulesOf(fs) != "R2" {
		t.Fatalf("leading zero: %+v", fs)
	}
	ref2 := func(p ...core.Word) rec {
		return recOf(core.Record{Kind: core.KindReference, Codec: core.RefExternalIssuer2Word << 3, Payload: p})
	}
	if fs := checkIssuerRef(ref2(1, 0x20, 0, 0, 0, 2, 7), core.RefExternalIssuer2Word); fs != nil {
		t.Fatalf("2-word: %+v", fs)
	}
	if fs := checkIssuerRef(ref2(0, 0x20, 0, 0, 0, 2, 7), core.RefExternalIssuer2Word); rulesOf(fs) != "R4" {
		t.Fatalf("2-word of a 1-word value: %+v", fs)
	}
}
