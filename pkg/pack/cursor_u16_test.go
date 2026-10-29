//ff:func feature=pack type=decoder control=sequence
//ff:what TestCursorU16 — little-endian, zero on overrun
package pack

import "testing"

func TestCursorU16(t *testing.T) {
	c := &cursor{b: []byte{1, 2, 3}}
	if c.u16() != 0x0201 || c.u16() != 0 || !c.bad {
		t.Fatal("u16")
	}
}
