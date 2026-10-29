//ff:func feature=pack type=writer control=sequence
//ff:what TestParseHEntry — fields come back from their little-endian places
package pack

import "testing"

func TestParseHEntry(t *testing.T) {
	b := make([]byte, hEntrySize)
	b[0], b[32], b[36], b[44], b[52] = 1, 2, 3, 4, encZstd
	e := parseHEntry(b)
	if e.key[0] != 1 || e.loc != (Loc{Obj: 2, Off: 3, Len: 4, Enc: encZstd}) {
		t.Fatalf("%+v", e)
	}
}
