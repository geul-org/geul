//ff:func feature=pack type=decoder control=sequence
//ff:what TestReadPaths — length-prefixed paths; an impossible count fails the cursor
package pack

import "testing"

func TestReadPaths(t *testing.T) {
	c := &cursor{b: []byte{1, 0, 0, 0, 2, 0, 'a', 'b'}}
	if p := readPaths(c); len(p) != 1 || p[0] != "ab" || c.bad {
		t.Fatal(p)
	}
	c = &cursor{b: []byte{9, 0, 0, 0}}
	if readPaths(c); !c.bad {
		t.Fatal("count")
	}
}
