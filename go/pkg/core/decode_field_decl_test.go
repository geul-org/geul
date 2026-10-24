//ff:func feature=core type=decoder control=iteration dimension=1
//ff:what TestDecodeFieldDecl — raw, record and expected-TAG fields round-trip with their DECLARE; other records, a missing DECLARE and bad layouts fail
package core

import (
	"slices"
	"testing"
)

func TestDecodeFieldDecl(t *testing.T) {
	for _, f := range []FieldDecl{{EdgeID: 2, Declare: 1, Width: 4}, {EdgeID: 3, Declare: 1}, {EdgeID: 4, Declare: 1, Expect: 0x00C0, HasExpect: true}} {
		got, err := DecodeFieldDecl(EncodeFieldDecl(f))
		if err != nil || got != f {
			t.Fatalf("%+v: %+v %v", f, got, err)
		}
	}
	id := localWords(1)
	for i, r := range []Record{
		EncodeUintValue(1),
		{Kind: KindEdge, Codec: EdgeField},
		{Kind: KindEdge, Codec: EdgeField, Payload: id},
		{Kind: KindEdge, Codec: EdgeField, Payload: slices.Concat(id, []Word{4})},
		{Kind: KindEdge, Codec: EdgeField, Payload: slices.Concat(id, id)},
		{Kind: KindEdge, Codec: EdgeField, Payload: slices.Concat(id, id, []Word{4, 0x00C0})},
	} {
		if _, err := DecodeFieldDecl(r); err == nil {
			t.Fatalf("bad %d: want error", i)
		}
	}
}
