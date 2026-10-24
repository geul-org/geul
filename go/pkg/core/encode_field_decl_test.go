//ff:func feature=core type=encoder control=sequence
//ff:what TestEncodeFieldDecl — edge/FIELD example words; an expected TAG on a raw field is dropped
package core

import (
	"slices"
	"testing"
)

func TestEncodeFieldDecl(t *testing.T) {
	if ws, _ := EncodeRecord(EncodeFieldDecl(FieldDecl{EdgeID: 7, Declare: 1, Expect: 0x0060, HasExpect: true})); !slices.Equal(ws, []Word{0x131C, 0x00E0, 7, 0x00E0, 1, 0, 0x0060}) {
		t.Fatalf("record: %04X", ws)
	}
	if ws, _ := EncodeRecord(EncodeFieldDecl(FieldDecl{EdgeID: 2, Declare: 1, Width: 4, Expect: 1, HasExpect: true})); !slices.Equal(ws, []Word{0x129C, 0x00E0, 2, 0x00E0, 1, 4}) {
		t.Fatalf("raw: %04X", ws)
	}
}
