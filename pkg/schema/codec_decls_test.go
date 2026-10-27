//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what TestCodecDecls — several CODECs on a record field pass; an unknown field, a repeated target and a second, wildcard or 0000 CODEC on a raw field and 0000 alone fail
package schema

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestCodecDecls(t *testing.T) {
	recs := func(rs ...core.Record) []rec {
		out := make([]rec, len(rs))
		for i, r := range rs {
			out[i] = recOf(r)
		}
		return out
	}
	field := func(id uint64, w core.Word) core.Record {
		return core.EncodeFieldDecl(core.FieldDecl{EdgeID: id, Declare: 1, Width: w})
	}
	codec := func(id, f uint64, t core.Word) core.Record {
		return core.EncodeCodecDecl(core.CodecDecl{EdgeID: id, Field: f, Target: t})
	}
	if fs := codecDecls(recs(field(2, 0), field(3, 1), codec(4, 2, 0x084A), codec(5, 2, 0x0780), codec(6, 3, 0x0802), codec(7, 2, 0x0000))); len(fs) != 0 {
		t.Fatalf("ok: %+v", fs)
	}
	for i, body := range [][]rec{
		recs(codec(4, 9, 0x084A)),
		recs(field(2, 0), codec(4, 2, 0x084A), codec(5, 2, 0x084A)),
		recs(field(3, 1), codec(4, 3, 0x0802), codec(5, 3, 0x0801)),
		recs(field(3, 1), codec(4, 3, 0x0F80)),
		recs(field(3, 1), codec(4, 3, 0x0802), codec(5, 3, 0x0000)),
		recs(field(2, 0), codec(4, 2, 0x0000)),
	} {
		if got := rulesOf(codecDecls(body)); got != "E5" {
			t.Fatalf("case %d: %q", i, got)
		}
	}
}
