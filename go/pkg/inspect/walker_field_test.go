//ff:func feature=inspect type=decoder control=iteration dimension=1
//ff:what TestWalkerField — raw words with ENUM names, the rest of the payload, records with their Value, expected TAG mismatch and short input marked bad
package inspect

import (
	"github.com/geul-org/geul/go/pkg/codecspec"
	"github.com/geul-org/geul/go/pkg/core"
	"testing"
)

func TestWalkerField(t *testing.T) {
	val, _ := core.EncodeRecord(core.EncodeUintValue(9))
	w := &walker{ws: append([]core.Word{1, 2}, val...)}
	w.field(codecspec.Field{Name: "code", Width: 1, Enums: map[string]string{"0x0001": "one"}})
	w.field(codecspec.Field{Name: "other", Width: 1})
	w.field(codecspec.Field{Name: "v", Expect: core.TargetWord(core.KindValue, core.CodecUnsignedInteger), HasExpect: true})
	if w.bad || w.out[0].Enum != "one" || w.out[1].Enum != "" || w.out[2].Data != (Value{Codec: core.CodecUnsignedInteger, Uint: 9}) {
		t.Fatalf("got %+v bad %v", w.out, w.bad)
	}
	r := &walker{ws: []core.Word{5, 6, 7}, pos: 1}
	r.field(codecspec.Field{Name: "rest", Width: int(core.FieldRest)})
	if r.bad || r.pos != 3 || len(r.out[0].Words) != 2 {
		t.Fatalf("rest: %+v", r)
	}
	r.field(codecspec.Field{Name: "empty", Width: int(core.FieldRest)})
	if r.bad || len(r.out) != 1 {
		t.Fatalf("empty rest: %+v", r)
	}
	for i, f := range []codecspec.Field{{Width: 3}, {Expect: 0x1018, HasExpect: true}} {
		w := &walker{ws: val}
		w.field(f)
		if !w.bad {
			t.Fatalf("bad %d: want bad", i)
		}
	}
}
