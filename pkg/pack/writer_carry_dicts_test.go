//ff:func feature=pack type=writer control=sequence
//ff:what TestWriterCarryDicts — renumbered previous DICTS; none on Fresh or for a new set
package pack

import "testing"

func TestWriterCarryDicts(t *testing.T) {
	prev := newReader(memStore{}, Options{}, &manifest{dicts: []dictEntry{{id: 1, loc: Loc{Obj: 4}}}}, "")
	w := &writer{prev: prev, remap: map[uint32]uint32{4: 0}}
	if d := w.carryDicts(); len(d) != 1 || d[0].loc.Obj != 0 {
		t.Fatal(d)
	}
	w.o.Fresh = true
	if d := w.carryDicts(); d != nil {
		t.Fatal(d)
	}
	if d := (&writer{}).carryDicts(); d != nil {
		t.Fatal("new set", d)
	}
}
