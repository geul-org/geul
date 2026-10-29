//ff:func feature=pack type=resolver control=sequence
//ff:what TestBlockChild — last first-key ≤ key; -1 before the first child
package pack

import "testing"

func TestBlockChild(t *testing.T) {
	b := block{kind: blkHInner, keys: []string{"b", "d"}}
	if b.child("a") != -1 || b.child("b") != 0 || b.child("c") != 0 || b.child("z") != 1 {
		t.Fatal("child")
	}
}
