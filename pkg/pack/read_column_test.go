//ff:func feature=pack type=decoder control=sequence
//ff:what TestReadColumn — n varints; a short column leaves the cursor failed
package pack

import "testing"

func TestReadColumn(t *testing.T) {
	c := &cursor{b: []byte{1, 0xAC, 2}}
	if v := readColumn(c, 2); v[0] != 1 || v[1] != 300 || c.bad {
		t.Fatal(v)
	}
	if readColumn(c, 1); !c.bad {
		t.Fatal("short")
	}
}
