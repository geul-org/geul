//ff:func feature=core type=encoder control=sequence
//ff:what TestEncodeIssuerPointer — a 1-word Issuer encodes as 00E8 + Issuer; a 3-word Issuer fails
package core

import (
	"slices"
	"testing"
)

func TestEncodeIssuerPointer(t *testing.T) {
	r, err := EncodeIssuerPointer([]Word{0x40})
	if err != nil {
		t.Fatal(err)
	}
	if ws, _ := EncodeRecord(r); !slices.Equal(ws, []Word{0x00B0, 0x40}) {
		t.Fatalf("%04X", ws)
	}
	if _, err := EncodeIssuerPointer([]Word{1, 2, 3}); err == nil {
		t.Fatal("3-word Issuer accepted")
	}
}
