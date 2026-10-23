//ff:func feature=core type=encoder control=sequence
//ff:what TestEncodeDeclare — the PUBLICATION example words of edge/DECLARE
package core

import (
	"slices"
	"testing"
)

func TestEncodeDeclare(t *testing.T) {
	ws, _ := EncodeRecord(EncodeDeclare(Declare{EdgeID: 1, Target: 0x1018, Fields: []uint64{2, 3, 4}}))
	want := []Word{0x149B, 0x00C0, 1, 0x1018, 0x00C0, 2, 0x00C0, 3, 0x00C0, 4}
	if !slices.Equal(ws, want) {
		t.Fatalf("got %04X", ws)
	}
}
