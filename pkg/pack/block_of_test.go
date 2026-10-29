//ff:func feature=pack type=writer control=sequence
//ff:what TestBlockOf — the block's keys and locations follow the entries
package pack

import "testing"

func TestBlockOf(t *testing.T) {
	b := blockOf(blkHInner, []streamKV{{Key: "a", Loc: Loc{Obj: 1}}, {Key: "b", Loc: Loc{Obj: 2}}})
	if b.kind != blkHInner || len(b.keys) != 2 || b.keys[1] != "b" || b.locs[0].Obj != 1 {
		t.Fatalf("%+v", b)
	}
	if b := blockOf(blkPLeaf, nil); len(b.keys) != 0 {
		t.Fatal("empty")
	}
}
