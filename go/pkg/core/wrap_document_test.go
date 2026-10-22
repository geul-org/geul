//ff:func feature=core type=encoder control=sequence
//ff:what TestWrapDocument — DOC-1WORD/2WORD selection by total length
package core

import "testing"

func TestWrapDocument(t *testing.T) {
	w, err := WrapDocument([]Word{0x180F})
	if err != nil || len(w) != 3 || w[0] != 0x1880 || w[1] != 3 {
		t.Fatalf("small: %04X %v", w, err)
	}
	w, err = WrapDocument(make([]Word, 0xFFFE))
	if err != nil || w[0] != 0x1901 || w[1] != 1 || w[2] != 1 {
		t.Fatalf("2-word: %04X %v", w[:3], err)
	}
}
