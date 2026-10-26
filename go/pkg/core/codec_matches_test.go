//ff:func feature=core type=helper control=iteration dimension=1
//ff:what TestCodecMatches — exact targets, any Reference, a Reference codec with any Quantifier, any Value; other Kinds and codecs do not match
package core

import "testing"

func TestCodecMatches(t *testing.T) {
	qid := RefWikidataQID << 3
	cases := []struct {
		t     Word
		k     Kind
		codec uint8
		want  bool
	}{
		{0x0020, KindReference, qid, true},
		{0x0020, KindReference, qid | 3, false},
		{0x0780, KindReference, qid | 4, true},
		{0x0780, KindValue, CodecUnsignedInteger, false},
		{0x07A0, KindReference, qid | 3, true},
		{0x07A0, KindReference, RefLocal << 3, false},
		{0x0781, KindReference, RefLocal<<3 | 1, true},
		{0x0781, KindReference, RefLocal << 3, false},
		{0x0F80, KindValue, CodecCoordinateEarth, true},
		{0x0F80, KindEdge, EdgeBind, false},
		{0x0F84, KindValue, CodecCoordinateEarth, false},
		{0x084A, KindValue, CodecUTF8LanguageTextEven | 1, false},
	}
	for i, c := range cases {
		if got := CodecMatches(c.t, c.k, c.codec); got != c.want {
			t.Fatalf("case %d: %v", i, got)
		}
	}
}
