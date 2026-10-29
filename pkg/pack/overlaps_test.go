//ff:func feature=pack type=helper control=sequence
//ff:what TestOverlaps — leaf keys in [lo, hi); internal children by their span; hi "" unbounded
package pack

import "testing"

func TestOverlaps(t *testing.T) {
	leaf := block{kind: blkHLeaf, keys: []string{"b", "d"}}
	if !overlaps(leaf, 0, "b", "c") || overlaps(leaf, 0, "c", "") || overlaps(leaf, 1, "a", "d") {
		t.Fatal("leaf")
	}
	in := block{kind: blkHInner, keys: []string{"b", "d"}}
	if !overlaps(in, 0, "c", "") || overlaps(in, 0, "d", "") || !overlaps(in, 1, "z", "") || overlaps(in, 1, "a", "c") {
		t.Fatal("internal")
	}
}
