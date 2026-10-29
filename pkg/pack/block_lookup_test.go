//ff:func feature=pack type=resolver control=sequence
//ff:what TestBlockLookup — exact hits only
package pack

import "testing"

func TestBlockLookup(t *testing.T) {
	b := block{kind: blkHLeaf, keys: []string{"a", "c"}, locs: []Loc{{Off: 1}, {Off: 2}}}
	if l, ok := b.lookup("c"); !ok || l.Off != 2 {
		t.Fatal(l)
	}
	if _, ok := b.lookup("b"); ok {
		t.Fatal("b")
	}
	if _, ok := b.lookup("d"); ok {
		t.Fatal("d")
	}
}
