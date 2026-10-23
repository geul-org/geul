//ff:func feature=core type=decoder control=sequence
//ff:what TestTargetWord — Kind and Codec bits round-trip through SplitTarget
package core

import (
	"testing"
)

func TestTargetWord(t *testing.T) {
	if w := TargetWord(KindEdge, EdgePublication); w != 0x1018 {
		t.Fatalf("got %04X", w)
	}
	if k, c := SplitTarget(TargetWord(KindMeta, 0x30)); k != KindMeta || c != 0x30 {
		t.Fatalf("got %v %02X", k, c)
	}
}
