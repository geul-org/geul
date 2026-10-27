//ff:func feature=core type=helper control=iteration dimension=1
//ff:what TestIsSlotMarker — WORDs 0x0000~0x0007 decode as one-word slot markers; LOCAL, a Codec 0000 Reference with payload, a Value and 0x0008 are not
package core

import "testing"

func TestIsSlotMarker(t *testing.T) {
	for w := Word(0); w < 8; w++ {
		r, n, err := DecodeRecord([]Word{w, 0x1234})
		if err != nil || n != 1 || !IsSlotMarker(r) {
			t.Fatalf("0x%04X: %+v %d %v", w, r, n, err)
		}
	}
	for _, r := range []Record{EncodeLocalRef(1), {Kind: KindReference, Codec: 1, Payload: []Word{1}}, {Kind: KindValue, Codec: 1}, {Kind: KindReference, Codec: RefLocal << 3}} {
		if IsSlotMarker(r) {
			t.Fatalf("%+v: want false", r)
		}
	}
}
