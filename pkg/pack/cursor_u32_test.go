//ff:func feature=pack type=decoder control=sequence
//ff:what TestCursorU32 — little-endian, zero on overrun
package pack

import "testing"

func TestCursorU32(t *testing.T) {
	c := &cursor{b: []byte{1, 2, 3, 4, 5}}
	if c.u32() != 0x04030201 || c.u32() != 0 || !c.bad {
		t.Fatal("u32")
	}
}
