//ff:func feature=core type=helper control=sequence
//ff:what TestWidthWords — every 2-bit width code maps to 1/2/4/8 words
package core

import "testing"

func TestWidthWords(t *testing.T) {
	if WidthWords(0) != 1 || WidthWords(1) != 2 || WidthWords(2) != 4 || WidthWords(3) != 8 {
		t.Fatal("width code table mismatch")
	}
}
