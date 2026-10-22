//ff:func feature=core type=encoder control=sequence
//ff:what TestEncodeID — fixed-width big-endian fields plus bad-width and overflow errors
package core

import "testing"

func TestEncodeID(t *testing.T) {
	if w, err := EncodeID(0x12345, 2); err != nil || len(w) != 2 || w[0] != 1 || w[1] != 0x2345 {
		t.Fatalf("2-word: %04X %v", w, err)
	}
	if w, err := EncodeID(5, 8); err != nil || len(w) != 8 || w[7] != 5 || w[0] != 0 {
		t.Fatalf("8-word: %04X %v", w, err)
	}
	if _, err := EncodeID(1, 3); err == nil {
		t.Fatal("width 3 must fail")
	}
	if _, err := EncodeID(0x10000, 1); err == nil {
		t.Fatal("overflow must fail")
	}
}
