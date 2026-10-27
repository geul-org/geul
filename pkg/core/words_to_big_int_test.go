//ff:func feature=core type=decoder control=sequence
//ff:what TestWordsToBigInt — sign from the top bit, empty and non-minimal input fail
package core

import "testing"

func TestWordsToBigInt(t *testing.T) {
	if n, err := wordsToBigInt([]Word{0xFFFF}); err != nil || n.Int64() != -1 {
		t.Fatalf("-1: %v %v", n, err)
	}
	if n, err := wordsToBigInt([]Word{0x0001, 0x0000, 0x0000, 0x0000, 0x0000}); err != nil || n.String() != "18446744073709551616" {
		t.Fatalf("2^64: %v %v", n, err)
	}
	if _, err := wordsToBigInt(nil); err == nil {
		t.Fatal("empty: want error")
	}
	if _, err := wordsToBigInt([]Word{0, 1}); err == nil {
		t.Fatal("non-minimal: want error")
	}
}
