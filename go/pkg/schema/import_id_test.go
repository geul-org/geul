//ff:func feature=schema type=rule control=sequence
//ff:what TestImportID — an IMPORT ID; too short and over 64 bits fail
package schema

import (
	"github.com/geul-org/geul/go/pkg/core"
	"testing"
)

func TestImportID(t *testing.T) {
	imp := func(p ...core.Word) rec {
		return recOf(core.Record{Kind: core.KindMeta, Codec: core.MetaImportDoc1Word, Payload: p})
	}
	if id, ok := importID(imp(1, 2, 1, 0, 0, 0, 2, 0, 0)); !ok || id != 0x10002 {
		t.Fatalf("ok: %d %v", id, ok)
	}
	if _, ok := importID(imp(1, 0, 0, 0, 2, 0, 0)); ok {
		t.Fatal("short: want false")
	}
	if _, ok := importID(imp(1, 1, 1, 1, 1, 1, 0, 0, 0, 2, 0, 0)); ok {
		t.Fatal("wide: want false")
	}
}
