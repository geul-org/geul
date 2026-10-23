//ff:func feature=core type=helper control=iteration dimension=1
//ff:what TestIssuerRefCodec — the four Issuer widths; other widths fail
package core

import "testing"

func TestIssuerRefCodec(t *testing.T) {
	for w, want := range map[int]uint8{1: 0x20, 2: 0x28, 4: 0x30, 16: 0x38} {
		if c, ok := issuerRefCodec(w); !ok || c != want {
			t.Fatalf("%d: %02X %v", w, c, ok)
		}
	}
	if _, ok := issuerRefCodec(3); ok {
		t.Fatal("3: want false")
	}
}
