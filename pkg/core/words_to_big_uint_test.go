//ff:func feature=core type=decoder control=sequence
//ff:what TestWordsToBigUint — 0, a value past uint64 and 0x8000 read back; an empty run and a leading 0x0000 fail
package core

import "testing"

func TestWordsToBigUint(t *testing.T) {
	if n, err := wordsToBigUint([]Word{0}); err != nil || n.Sign() != 0 {
		t.Fatalf("zero: %v %v", n, err)
	}
	if n, err := wordsToBigUint([]Word{1, 0, 0, 0, 0}); err != nil || n.String() != "18446744073709551616" {
		t.Fatalf("2^64: %v %v", n, err)
	}
	if n, err := wordsToBigUint([]Word{0x8000}); err != nil || n.Int64() != 0x8000 {
		t.Fatalf("0x8000: %v %v", n, err)
	}
	if _, err := wordsToBigUint(nil); err == nil {
		t.Fatal("empty: want error")
	}
	if _, err := wordsToBigUint([]Word{0, 1}); err == nil {
		t.Fatal("leading zero: want error")
	}
}
