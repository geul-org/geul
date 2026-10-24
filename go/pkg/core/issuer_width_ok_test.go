//ff:func feature=core type=helper control=iteration dimension=1
//ff:what TestIssuerWidthOK — 1, 2, 4 and 16 words pass; other widths fail
package core

import "testing"

func TestIssuerWidthOK(t *testing.T) {
	for w := 0; w <= 17; w++ {
		want := w == 1 || w == 2 || w == 4 || w == 16
		if issuerWidthOK(w) != want {
			t.Fatalf("%d: want %v", w, want)
		}
	}
}
