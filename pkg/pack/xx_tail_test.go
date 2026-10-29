//ff:func feature=pack type=helper control=sequence
//ff:what TestXxTail — 8-, 4- and 1-byte steps (checked through known XXH64 values)
package pack

import "testing"

func TestXxTail(t *testing.T) {
	if xxTail(xxP5+3, []byte("abc")) != 0x44bc2cf5ad770999 {
		t.Fatal("abc")
	}
	b := []byte("0123456789abc")
	if xxTail(xxP5+13, b) != xxh64(b) {
		t.Fatal("13")
	}
}
