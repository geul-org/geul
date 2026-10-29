//ff:func feature=pack type=decoder control=iteration dimension=1
//ff:what TestCursorUvarint — shortest LEB128 read; padding zero bytes, overflow, truncation and a failed cursor are rejected
package pack

import "testing"

func TestCursorUvarint(t *testing.T) {
	c := &cursor{b: []byte{0x00, 0xAC, 0x02}}
	if c.uvarint() != 0 || c.uvarint() != 300 || c.bad || len(c.b) != 0 {
		t.Fatal("valid")
	}
	for _, b := range [][]byte{{0x80, 0x00}, {0x81}, {0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x02}} {
		c := &cursor{b: b}
		if c.uvarint() != 0 || !c.bad {
			t.Fatalf("%x accepted", b)
		}
	}
	c = &cursor{b: []byte{1}, bad: true}
	if c.uvarint() != 0 {
		t.Fatal("failed cursor")
	}
	max := &cursor{b: []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x01}}
	if max.uvarint() != ^uint64(0) || max.bad {
		t.Fatal("2^64-1")
	}
}
