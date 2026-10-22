//ff:func feature=core type=decoder control=sequence
//ff:what TestDecodeLength — direct, extension, non-canonical and truncated Length codes
package core

import "testing"

func TestDecodeLength(t *testing.T) {
	if total, head, err := decodeLength([]Word{0}, 3); err != nil || total != 4 || head != 1 {
		t.Fatalf("direct: %d %d %v", total, head, err)
	}
	if total, head, err := decodeLength([]Word{0, 15}, 0xE); err != nil || total != 15 || head != 2 {
		t.Fatalf("ext1: %d %d %v", total, head, err)
	}
	if _, _, err := decodeLength([]Word{0, 14}, 0xE); err == nil {
		t.Fatal("ext1 below 15 must fail")
	}
	if _, _, err := decodeLength([]Word{0}, 0xE); err == nil {
		t.Fatal("ext1 truncated must fail")
	}
	if total, head, err := decodeLength([]Word{0, 1, 0}, 0xF); err != nil || total != 0x10000 || head != 3 {
		t.Fatalf("ext2: %d %d %v", total, head, err)
	}
	if _, _, err := decodeLength([]Word{0, 0, 0xFFFF}, 0xF); err == nil {
		t.Fatal("ext2 below 65536 must fail")
	}
	if _, _, err := decodeLength([]Word{0, 1}, 0xF); err == nil {
		t.Fatal("ext2 truncated must fail")
	}
}
