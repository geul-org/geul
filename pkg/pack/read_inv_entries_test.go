//ff:func feature=pack type=decoder control=sequence
//ff:what TestReadInvEntries — 32-byte entries; an impossible count fails the cursor
package pack

import "testing"

func TestReadInvEntries(t *testing.T) {
	b := append([]byte{1, 0, 0, 0, 0, 0, 0, 0, invP, 1}, make([]byte, 30)...)
	c := &cursor{b: b}
	if e := readInvEntries(c); len(e) != 1 || e[0].kind != invP || e[0].enc != 1 || c.bad {
		t.Fatal(e)
	}
	c = &cursor{b: []byte{1, 0, 0, 0, 0, 0, 0, 0}}
	if readInvEntries(c); !c.bad {
		t.Fatal("count")
	}
}
