//ff:func feature=core type=encoder control=sequence
//ff:what TestTagWord — bit placement of Kind, Length code and Codec
package core

import "testing"

func TestTagWord(t *testing.T) {
	if got := tagWord(KindMeta, 1, 0x00); got != 0x1880 {
		t.Fatalf("DOC-1WORD tag: got 0x%04X", uint16(got))
	}
	if got := tagWord(KindReference, 0, 0x7F); got != 0x007F {
		t.Fatalf("reference tag: got 0x%04X", uint16(got))
	}
}
