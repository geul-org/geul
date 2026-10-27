//ff:func feature=core type=encoder control=sequence
//ff:what TestEncodeFieldDecl — edge/FIELD example words for a record and a raw field
package core

import (
	"slices"
	"testing"
)

func TestEncodeFieldDecl(t *testing.T) {
	if ws, _ := EncodeRecord(EncodeFieldDecl(FieldDecl{EdgeID: 7, Declare: 1})); !slices.Equal(ws, []Word{0x129C, 0x0088, 7, 0x0088, 1, 0}) {
		t.Fatalf("record: %04X", ws)
	}
	if ws, _ := EncodeRecord(EncodeFieldDecl(FieldDecl{EdgeID: 2, Declare: 1, Width: 4})); !slices.Equal(ws, []Word{0x129C, 0x0088, 2, 0x0088, 1, 4}) {
		t.Fatalf("raw: %04X", ws)
	}
}
