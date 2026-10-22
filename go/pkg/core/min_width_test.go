//ff:func feature=core type=helper control=sequence
//ff:what TestMinWidth — 1/2/4-word boundaries
package core

import "testing"

func TestMinWidth(t *testing.T) {
	if MinWidth(0xFFFF) != 1 || MinWidth(0x10000) != 2 || MinWidth(0xFFFFFFFF) != 2 || MinWidth(1<<32) != 4 {
		t.Fatal("MinWidth boundary mismatch")
	}
}
