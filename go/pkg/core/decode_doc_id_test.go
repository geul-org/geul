//ff:func feature=core type=decoder control=iteration dimension=1
//ff:what TestDecodeDocID — round trip; other records, junk, trailing words, LOCAL, the old EXTERNAL-ISSUER form, a version, document 0 and a non-Identified Quantifier fail
package core

import (
	"slices"
	"testing"
)

func TestDecodeDocID(t *testing.T) {
	r, _ := EncodeDocID([]Word{1, 0x20}, []Word{0xffff, 0, 0, 1})
	if i, d, err := DecodeDocID(r); err != nil || !slices.Equal(i, []Word{1, 0x20}) || !slices.Equal(d, []Word{0xffff, 0, 0, 1}) {
		t.Fatalf("got %v %v %v", i, d, err)
	}
	meta := func(p ...Word) Record { return Record{Kind: KindMeta, Codec: MetaDocID, Payload: p} }
	for i, b := range []Record{
		EncodeUintValue(1),
		meta(0, 0, 0, 0, 1),
		meta(append(slices.Clone(r.Payload), 9)...),
		meta(0x0088, 1),
		meta(0x0360, 0, 0xffff, 0, 0, 1, 0),
		meta(0x03B8, 0, 0xffff, 0, 0, 1, 0, 1),
		meta(0x03B8, 0, 0, 0, 0, 0, 0, 0),
		meta(0x03B9, 0, 0xffff, 0, 0, 1, 0, 0),
	} {
		if _, _, err := DecodeDocID(b); err == nil {
			t.Fatalf("bad %d: want error", i)
		}
	}
}
