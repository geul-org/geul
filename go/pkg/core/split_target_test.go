//ff:func feature=core type=decoder control=sequence
//ff:what TestSplitTarget — a Reference target keeps its CCCC<<3 codec
package core

import (
	"testing"
)

func TestSplitTarget(t *testing.T) {
	if k, c := SplitTarget(0x0060); k != KindReference || c != 0x60 {
		t.Fatalf("got %v %02X", k, c)
	}
}
