//ff:func feature=core type=decoder control=iteration dimension=1
//ff:what TestDecodeDocPointer — an encoded document version round-trips; document 0 decodes; other codecs and widths fail
package core

import (
	"slices"
	"testing"
)

func TestDecodeDocPointer(t *testing.T) {
	r, _ := EncodeDocPointer([]Word{1, 0x40}, []Word{0xFFFF, 0, 0, 1}, []Word{0, 2})
	iss, doc, ver, err := DecodeDocPointer(r)
	if err != nil || !slices.Equal(iss, []Word{1, 0x40}) || !slices.Equal(doc, []Word{0xFFFF, 0, 0, 1}) || !slices.Equal(ver, []Word{0, 2}) {
		t.Fatalf("%v %v %v %v", iss, doc, ver, err)
	}
	if _, d, _, err := DecodeDocPointer(Record{Kind: KindReference, Codec: RefDocPointer << 3, Payload: []Word{0, 0, 0, 0, 0, 0, 1}}); err != nil || d[3] != 0 {
		t.Fatalf("document 0: %v", err)
	}
	for _, bad := range []Record{
		{Kind: KindReference, Codec: RefDocPointer<<3 | 4, Payload: []Word{0, 0xFFFF, 0, 0, 1, 0, 1}},
		{Kind: KindReference, Codec: RefDocPointer << 3, Payload: []Word{0, 0, 0, 0xFFFF, 0, 0, 1, 0, 1}},
		{Kind: KindReference, Codec: RefIssuerPointer << 3, Payload: []Word{0, 0xFFFF, 0, 0, 1, 0, 1}},
	} {
		if _, _, _, err := DecodeDocPointer(bad); err == nil {
			t.Fatalf("%+v accepted", bad)
		}
	}
}
