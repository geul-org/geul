//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what TestHEntryLess — key first, then object, then offset
package pack

import "testing"

func TestHEntryLess(t *testing.T) {
	at := func(k byte, obj uint32, off uint64) hEntry {
		e := hEntry{loc: Loc{Obj: obj, Off: off}}
		e.key[0] = k
		return e
	}
	for _, c := range []struct {
		a, b hEntry
		want bool
	}{
		{at(1, 9, 9), at(2, 0, 0), true}, {at(2, 0, 0), at(1, 9, 9), false},
		{at(1, 1, 9), at(1, 2, 0), true}, {at(1, 2, 0), at(1, 1, 9), false},
		{at(1, 1, 1), at(1, 1, 2), true}, {at(1, 1, 2), at(1, 1, 1), false}, {at(1, 1, 1), at(1, 1, 1), false},
	} {
		if hEntryLess(c.a, c.b) != c.want {
			t.Errorf("%v %v", c.a.loc, c.b.loc)
		}
	}
}
