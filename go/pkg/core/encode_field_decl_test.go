//ff:func feature=core type=encoder control=sequence
//ff:what TestEncodeFieldDecl — edge/FIELD example words; an expected TAG on a raw field is dropped
package core

import (
	"slices"
	"testing"
)

func TestEncodeFieldDecl(t *testing.T) {
	if ws, _ := EncodeRecord(EncodeFieldDecl(FieldDecl{EdgeID: 7, Declare: 1, Expect: 0x0040, HasExpect: true})); !slices.Equal(ws, []Word{0x131C, 0x00C0, 7, 0x00C0, 1, 0, 0x0040}) {
		t.Fatalf("record: %04X", ws)
	}
	if ws, _ := EncodeRecord(EncodeFieldDecl(FieldDecl{EdgeID: 2, Declare: 1, Width: 4, Expect: 1, HasExpect: true})); !slices.Equal(ws, []Word{0x129C, 0x00C0, 2, 0x00C0, 1, 4}) {
		t.Fatalf("raw: %04X", ws)
	}
}
