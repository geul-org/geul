//ff:func feature=pack type=writer control=sequence
//ff:what TestWriterMapLoc — renumbered with a remap, unchanged without
package pack

import "testing"

func TestWriterMapLoc(t *testing.T) {
	if (&writer{}).mapLoc(Loc{Obj: 3}).Obj != 3 || (&writer{remap: map[uint32]uint32{3: 1}}).mapLoc(Loc{Obj: 3}).Obj != 1 {
		t.Fatal("mapLoc")
	}
}
