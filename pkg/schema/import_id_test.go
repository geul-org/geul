//ff:func feature=schema type=rule control=sequence
//ff:what TestImportID — an IMPORT ID; too short and over 64 bits fail
package schema

import (
	"github.com/geul-org/gogeul/pkg/core"
	"testing"
)

func TestImportID(t *testing.T) {
	if id, ok := importID(importDoc(0, 1, 2)); !ok || id != 0x10002 {
		t.Fatalf("ok: %d %v", id, ok)
	}
	if _, ok := importID(importDoc(0)); ok {
		t.Fatal("short: want false")
	}
	if _, ok := importID(importDoc(0, 1, 1, 1, 1, 1)); ok {
		t.Fatal("wide: want false")
	}
	if _, ok := importID(recOf(core.Record{Kind: core.KindMeta, Codec: core.MetaImportDoc, Payload: []core.Word{1, 1, 0, 0, 0, 2, 0, 0}})); ok {
		t.Fatal("no DOC-POINTER: want false")
	}
}
