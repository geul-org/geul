//ff:func feature=pack type=decoder control=sequence
//ff:what TestCursorTake — bytes come off the front; an overrun fails and stays failed
package pack

import "testing"

func TestCursorTake(t *testing.T) {
	c := &cursor{b: []byte{1, 2, 3}}
	if b := c.take(2); len(b) != 2 || b[1] != 2 || c.bad {
		t.Fatal(b)
	}
	if b := c.take(2); b != nil || !c.bad {
		t.Fatal("overrun")
	}
	if b := c.take(0); b != nil {
		t.Fatal("sticky")
	}
}
