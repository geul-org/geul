//ff:func feature=pack type=writer control=sequence
//ff:what TestWriterNewState — previous state renumbered, deletions removed, puts (and shared h/ puts) placed
package pack

import "testing"

func TestWriterNewState(t *testing.T) {
	p := &put{path: "p", loc: Loc{Obj: 5}}
	w := &writer{base: map[string]Loc{"a": {Obj: 2}, "d": {}, "p": {Obj: 2}}, remap: map[uint32]uint32{2: 0},
		dels: []string{"d"}, puts: []*put{p, {path: "h", via: p}}}
	st := w.newState()
	if len(st) != 3 || st["a"].Obj != 0 || st["p"].Obj != 5 || st["h"].Obj != 5 {
		t.Fatal(st)
	}
}
