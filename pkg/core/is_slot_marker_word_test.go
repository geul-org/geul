//ff:func feature=core type=helper control=iteration dimension=1
//ff:what TestIsSlotMarkerWord — 0x0000~0x0007 are slot marker WORDs; 0x0008 (LOCAL), 0x0780 (wildcard) and 0x0801 are not
package core

import "testing"

func TestIsSlotMarkerWord(t *testing.T) {
	for w, want := range map[Word]bool{Absent: true, UnknownValue: true, NoValue: true, 0x0007: true, 0x0008: false, 0x0780: false, 0x0801: false} {
		if IsSlotMarkerWord(w) != want {
			t.Fatalf("0x%04X: want %v", w, want)
		}
	}
}
