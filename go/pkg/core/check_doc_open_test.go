//ff:func feature=core type=decoder control=iteration dimension=1
//ff:what TestCheckDocOpen — DOC kind/codec, payload width, length match and canonical width checks
package core

import "testing"

func TestCheckDocOpen(t *testing.T) {
	ok := Record{Kind: KindMeta, Codec: MetaDoc1Word, Payload: []Word{3}}
	if err := checkDocOpen(ok, 3); err != nil {
		t.Fatal(err)
	}
	bad := []struct {
		r Record
		n int
	}{
		{Record{Kind: KindEdge, Codec: 0}, 1},
		{Record{Kind: KindMeta, Codec: MetaDoc2Word, Payload: []Word{3}}, 3},
		{Record{Kind: KindMeta, Codec: MetaDoc1Word, Payload: []Word{4}}, 3},
		{Record{Kind: KindMeta, Codec: MetaDoc2Word, Payload: []Word{0, 4}}, 4},
	}
	for _, b := range bad {
		if err := checkDocOpen(b.r, b.n); err == nil {
			t.Fatalf("expected error for %+v", b.r)
		}
	}
}
