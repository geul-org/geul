//ff:func feature=pack type=encoder control=sequence
//ff:what TestEncodeDicts — u32 count then 52-byte entries
package pack

import "testing"

func TestEncodeDicts(t *testing.T) {
	b := encodeDicts([]dictEntry{{id: 40000, loc: Loc{Obj: 1, Off: 64, Len: 9}}})
	if len(b) != 4+52 || b[8] != 1 || b[12] != 64 || b[20] != 9 {
		t.Fatalf("%x", b)
	}
}
