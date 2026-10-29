//ff:func feature=pack type=encoder control=sequence
//ff:what TestOffCode — deletion marker 1, contiguous 0, otherwise offset + 1
package pack

import "testing"

func TestOffCode(t *testing.T) {
	ls := []Loc{{Obj: 0, Off: 64, Len: 4}, {Obj: 0, Off: 68, Len: 1}, {Enc: encGone}, {Obj: 0, Off: 69}}
	if offCode(ls, 0) != 65 || offCode(ls, 1) != 0 || offCode(ls, 2) != 1 || offCode(ls, 3) != 70 {
		t.Fatal("offCode")
	}
}
