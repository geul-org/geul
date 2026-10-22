//ff:func feature=core type=decoder control=sequence
//ff:what TestDecodeID — round trip plus bad width, truncation and above-64-bit errors
package core

import "testing"

func TestDecodeID(t *testing.T) {
	w, _ := EncodeID(1<<40, 4)
	if v, err := DecodeID(w, 4); err != nil || v != 1<<40 {
		t.Fatalf("got %d %v", v, err)
	}
	if _, err := DecodeID(w, 3); err == nil {
		t.Fatal("width 3 must fail")
	}
	if _, err := DecodeID([]Word{1}, 2); err == nil {
		t.Fatal("truncated must fail")
	}
	if _, err := DecodeID([]Word{1, 0, 0, 0, 0, 0, 0, 0}, 8); err == nil {
		t.Fatal("value above 64 bits must fail")
	}
}
