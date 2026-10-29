//ff:func feature=pack type=helper control=sequence
//ff:what TestIsPKind — kinds 1 and 2 are P
package pack

import "testing"

func TestIsPKind(t *testing.T) {
	if !isPKind(blkPLeaf) || !isPKind(blkPInner) || isPKind(blkHLeaf) || isPKind(blkHInner) {
		t.Fatal("isPKind")
	}
}
