//ff:func feature=core type=helper control=iteration dimension=1
//ff:what TestWidthCode — inverse of WidthWords plus the invalid-width error
package core

import "testing"

func TestWidthCode(t *testing.T) {
	for code := uint8(0); code < 4; code++ {
		if got, err := WidthCode(WidthWords(code)); err != nil || got != code {
			t.Fatalf("code %d: got %d %v", code, got, err)
		}
	}
	if _, err := WidthCode(3); err == nil {
		t.Fatal("width 3 must fail")
	}
}
