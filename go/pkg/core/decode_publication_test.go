//ff:func feature=core type=decoder control=iteration dimension=1
//ff:what TestDecodePublication — the three codecs decode; wrong kind, bad Edge Id, wrong length and version 0 fail
package core

import "testing"

func TestDecodePublication(t *testing.T) {
	id, _ := EncodeRecord(EncodeLocalRef(3))
	edge := func(c uint8, rest ...Word) Record {
		return Record{Kind: KindEdge, Codec: c, Payload: append(append([]Word{}, id...), rest...)}
	}
	pub := edge(EdgePublication, append([]Word{0, 0, 0, 5, 0, 2}, make([]Word, 16)...)...)
	if p, err := DecodePublication(pub); err != nil || p.EdgeID != 3 || p.Doc[3] != 5 || p.Version != 2 {
		t.Fatalf("publication: %+v %v", p, err)
	}
	if p, err := DecodePublication(edge(EdgeTombstone, 0, 0, 0, 5, 0, 1, 9)); err != nil || p.Reason != 9 {
		t.Fatalf("tombstone: %+v %v", p, err)
	}
	bad := []Record{
		{Kind: KindValue, Codec: EdgePublication},
		{Kind: KindEdge, Codec: EdgeDeprecation},
		{Kind: KindEdge, Codec: EdgeDeprecation, Payload: []Word{0x0882, 1, 0, 0, 0, 5, 0, 1}},
		edge(EdgeDeprecation, 0, 0, 0, 5, 0),
		edge(EdgeDeprecation, 0, 0, 0, 5, 0, 0),
	}
	for i, r := range bad {
		if _, err := DecodePublication(r); err == nil {
			t.Fatalf("bad %d: want error", i)
		}
	}
}
