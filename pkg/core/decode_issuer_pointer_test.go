//ff:func feature=core type=decoder control=iteration dimension=1
//ff:what TestDecodeIssuerPointer — an encoded Issuer round-trips; other codecs, Quantifiers and widths fail
package core

import (
	"slices"
	"testing"
)

func TestDecodeIssuerPointer(t *testing.T) {
	r, _ := EncodeIssuerPointer([]Word{1, 0x40})
	if iss, err := DecodeIssuerPointer(r); err != nil || !slices.Equal(iss, []Word{1, 0x40}) {
		t.Fatalf("%v %v", iss, err)
	}
	for _, bad := range []Record{
		{Kind: KindReference, Codec: RefIssuerPointer<<3 | 4, Payload: []Word{0x40}},
		{Kind: KindReference, Codec: RefIssuerPointer << 3, Payload: []Word{1, 2, 3}},
		{Kind: KindReference, Codec: RefDocPointer << 3, Payload: []Word{0x40}},
		{Kind: KindValue, Codec: RefIssuerPointer << 3, Payload: []Word{0x40}},
	} {
		if _, err := DecodeIssuerPointer(bad); err == nil {
			t.Fatalf("%+v accepted", bad)
		}
	}
}
