//ff:func feature=schema type=rule control=sequence
//ff:what TestAbsentOnly — 0000 with another target passes; 0000 alone is E5; non-CODEC records are skipped
package schema

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestAbsentOnly(t *testing.T) {
	codec := func(id, f uint64, tg core.Word) core.Record {
		return core.EncodeCodecDecl(core.CodecDecl{EdgeID: id, Field: f, Target: tg})
	}
	ok := []rec{recOf(core.EncodeLocalRef(1)), recOf(codec(4, 2, 0x0861)), recOf(codec(5, 2, 0x0000))}
	if fs := absentOnly(ok); len(fs) != 0 {
		t.Fatalf("ok: %+v", fs)
	}
	if got := rulesOf(absentOnly([]rec{recOf(codec(4, 2, 0x0000)), recOf(codec(5, 3, 0x0861))})); got != "E5" {
		t.Fatalf("alone: %q", got)
	}
}
