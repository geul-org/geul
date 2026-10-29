//ff:func feature=pack type=helper control=sequence
//ff:what TestContiguous — same object, next offset, previous not a deletion marker
package pack

import "testing"

func TestContiguous(t *testing.T) {
	p := Loc{Obj: 1, Off: 64, Len: 10}
	if !contiguous(p, Loc{Obj: 1, Off: 74}) || contiguous(p, Loc{Obj: 2, Off: 74}) || contiguous(p, Loc{Obj: 1, Off: 75}) || contiguous(Loc{Enc: encGone}, Loc{}) {
		t.Fatal("contiguous")
	}
}
