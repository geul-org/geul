//ff:func feature=pack type=helper control=sequence
//ff:what TestIsLeaf — kinds 1 and 3 are leaves
package pack

import "testing"

func TestIsLeaf(t *testing.T) {
	if !isLeaf(blkPLeaf) || !isLeaf(blkHLeaf) || isLeaf(blkPInner) || isLeaf(blkHInner) {
		t.Fatal("isLeaf")
	}
}
