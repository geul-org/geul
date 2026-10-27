//ff:func feature=core type=encoder control=sequence
//ff:what TestUintWords — minimal word forms of 0, one-word and multi-word values
package core

import "testing"

func TestUintWords(t *testing.T) {
	if w := uintWords(0); len(w) != 1 || w[0] != 0 {
		t.Fatalf("0: %v", w)
	}
	if w := uintWords(0xFFFF); len(w) != 1 {
		t.Fatalf("0xFFFF: %v", w)
	}
	if w := uintWords(0x10000); len(w) != 2 || w[0] != 1 || w[1] != 0 {
		t.Fatalf("0x10000: %v", w)
	}
}
