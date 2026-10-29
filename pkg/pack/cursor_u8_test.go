//ff:func feature=pack type=decoder control=sequence
//ff:what TestCursorU8 — one byte, then zero on overrun
package pack

import "testing"

func TestCursorU8(t *testing.T) {
	c := &cursor{b: []byte{7}}
	if c.u8() != 7 || c.u8() != 0 || !c.bad {
		t.Fatal("u8")
	}
}
