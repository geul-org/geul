//ff:func feature=core type=decoder control=sequence
//ff:what TestReadWords — odd-length byte input is rejected (words are 2-byte aligned)
package core

import "testing"

func TestReadWords(t *testing.T) {
	if _, err := ReadWords([]byte{0x12, 0x34, 0x56}); err == nil {
		t.Fatal("expected error for odd-length input")
	}
}
