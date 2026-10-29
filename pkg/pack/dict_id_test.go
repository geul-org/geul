//ff:func feature=pack type=decoder control=sequence
//ff:what TestDictID — a formatted dictionary's ID; raw content and short bytes have none
package pack

import "testing"

func TestDictID(t *testing.T) {
	if id, ok := dictID([]byte{0x37, 0xa4, 0x30, 0xec, 0x00, 0x80, 0, 0, 1}); !ok || id != 0x8000 {
		t.Fatal(id, ok)
	}
	if _, ok := dictID([]byte("raw content dictionary")); ok {
		t.Fatal("raw")
	}
	if _, ok := dictID([]byte{0x37, 0xa4, 0x30, 0xec}); ok {
		t.Fatal("short")
	}
}
