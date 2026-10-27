//ff:func feature=schema type=rule control=sequence
//ff:what TestCheckPointer — short payload, wider codec than needed, missing IMPORT-DOC, target 0, ok
package schema

import (
	"github.com/geul-org/gogeul/pkg/core"
	"testing"
)

func TestCheckPointer(t *testing.T) {
	ref := func(c uint8, p ...core.Word) rec {
		return recOf(core.Record{Kind: core.KindReference, Codec: c << 3, Payload: p})
	}
	imports := map[uint64]bool{1: true}
	if fs := checkPointer(ref(0, 1, 5), 0, imports); fs != nil {
		t.Fatalf("ok: %+v", fs)
	}
	if fs := checkPointer(ref(0, 1), 0, imports); rulesOf(fs) != "R3" {
		t.Fatalf("short: %+v", fs)
	}
	if fs := checkPointer(ref(1, 0, 1, 5), 1, imports); rulesOf(fs) != "R3" {
		t.Fatalf("wide: %+v", fs)
	}
	if fs := checkPointer(ref(0, 2, 0, 5), 0, imports); rulesOf(fs) != "R3,R2" {
		t.Fatalf("missing + target: %+v", fs)
	}
	if fs := checkPointer(ref(0, 1, 0), 0, imports); rulesOf(fs) != "R3" {
		t.Fatalf("document itself: %+v", fs)
	}
}
