//ff:func feature=core type=helper control=iteration dimension=1
//ff:what TestSlotMarkerName — 0 ABSENT, 1 UNKNOWN-VALUE, 2 NO-VALUE, 3~7 reserved (""), only the low 3 bits read
package core

import "testing"

func TestSlotMarkerName(t *testing.T) {
	want := []string{"ABSENT", "UNKNOWN-VALUE", "NO-VALUE", "", "", "", "", "", "ABSENT", "UNKNOWN-VALUE"}
	for m, name := range want {
		if got := SlotMarkerName(uint8(m)); got != name {
			t.Fatalf("SlotMarkerName(%d)=%q want %q", m, got, name)
		}
	}
}
