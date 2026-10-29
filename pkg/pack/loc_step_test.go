//ff:func feature=pack type=decoder control=iteration dimension=1
//ff:what TestLocStep — offsets, contiguity, deletion markers and their invalid forms
package pack

import (
	"math"
	"testing"
)

func TestLocStep(t *testing.T) {
	prev := []Loc{{Obj: 1, Off: 64, Len: 5}}
	if l, ok := locStep(prev, 1, 1, 0, 3, 0, true); !ok || l.Off != 69 {
		t.Fatal("contiguous", l)
	}
	if l, ok := locStep(nil, 0, 1, 65, 3, 1, true); !ok || l.Off != 64 || l.Enc != 1 {
		t.Fatal("offset", l)
	}
	if l, ok := locStep(nil, 0, 0, 1, 0, encGone, true); !ok || l.Enc != encGone {
		t.Fatal("marker")
	}
	bad := []struct {
		locs        []Loc
		j           int
		obj, off, n uint64
		enc         byte
		leaf        bool
	}{
		{nil, 0, math.MaxUint32 + 1, 1, 0, 0, true},
		{nil, 0, 0, 1, 0, encGone, false},
		{nil, 0, 0, 2, 0, encGone, true},
		{nil, 0, 0, 1, 0, 2, true},
		{nil, 0, 0, 0, 0, 0, true},
		{[]Loc{{Enc: encGone}}, 1, 0, 0, 0, 0, true},
		{prev, 1, 2, 0, 0, 0, true},
	}
	for i, c := range bad {
		if _, ok := locStep(c.locs, c.j, c.obj, c.off, c.n, c.enc, c.leaf); ok {
			t.Fatal(i)
		}
	}
}
