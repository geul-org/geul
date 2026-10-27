//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what TestCheckSpecVersion — 0x0002 passes; another version warns; empty, non-minimal and an embedded Value fail M4
package schema

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestCheckSpecVersion(t *testing.T) {
	sv := func(p ...core.Word) rec {
		return recOf(core.Record{Kind: core.KindMeta, Codec: core.MetaSpecVersion, Payload: p})
	}
	if checkSpecVersion(sv(2)) != nil {
		t.Fatal("0x0002")
	}
	if fs := checkSpecVersion(sv(3)); rulesOf(fs) != "M4" || fs[0].Severity != Warn {
		t.Fatalf("0x0003: %+v", fs)
	}
	for _, r := range []rec{sv(), sv(0, 2), recOf(metaOf(core.MetaSpecVersion, core.EncodeBytesValue(0x40, []byte("geul-t002"))))} {
		if fs := checkSpecVersion(r); rulesOf(fs) != "M4" || fs[0].Severity != Error {
			t.Fatalf("%04X: want M4", r.R.Payload)
		}
	}
}
