//ff:func feature=pack type=writer control=sequence
//ff:what TestAppendHEntry — an entry is hEntrySize bytes and parses back
package pack

import "testing"

func TestAppendHEntry(t *testing.T) {
	e := hEntry{loc: Loc{Obj: 3, Off: 1 << 40, Len: 77, Enc: encZstd}}
	e.key[0], e.key[31] = 0xAB, 0xCD
	b := appendHEntry([]byte{9}, e)
	if len(b) != 1+hEntrySize || parseHEntry(b[1:]) != e {
		t.Fatalf("%x", b)
	}
}
