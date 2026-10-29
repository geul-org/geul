//ff:func feature=pack type=decoder control=sequence
//ff:what TestCursorU64 — little-endian, zero on overrun
package pack

import "testing"

func TestCursorU64(t *testing.T) {
	c := &cursor{b: []byte{1, 0, 0, 0, 0, 0, 0, 0x80, 9}}
	if c.u64() != 0x8000000000000001 || c.u64() != 0 || !c.bad {
		t.Fatal("u64")
	}
}
