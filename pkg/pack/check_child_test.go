//ff:func feature=pack type=rule control=iteration dimension=1
//ff:what TestCheckChild — same space and the parent's key as first key; otherwise ErrCorrupt
package pack

import (
	"errors"
	"testing"
)

func TestCheckChild(t *testing.T) {
	parent := block{kind: blkPInner, keys: []string{"a", "c"}}
	if checkChild(parent, 1, block{kind: blkPLeaf, keys: []string{"c", "d"}}) != nil {
		t.Fatal("ok")
	}
	for _, c := range []block{{kind: blkHLeaf, keys: []string{"c"}}, {kind: blkPLeaf}, {kind: blkPLeaf, keys: []string{"d"}}} {
		if !errors.Is(checkChild(parent, 1, c), ErrCorrupt) {
			t.Fatal(c)
		}
	}
}
