//ff:func feature=pack type=decoder control=sequence
//ff:what TestCursorHash — 32 bytes, zero hash on overrun
package pack

import "testing"

func TestCursorHash(t *testing.T) {
	b := make([]byte, 33)
	b[0], b[31] = 1, 2
	c := &cursor{b: b}
	if h := c.hash(); h[0] != 1 || h[31] != 2 {
		t.Fatal(h)
	}
	if h := c.hash(); h != [32]byte{} || !c.bad {
		t.Fatal("overrun")
	}
}
