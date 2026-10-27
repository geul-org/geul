//ff:func feature=core type=decoder control=iteration dimension=1
//ff:what TestWordsToInt — round trip of intWords plus empty, removable sign word and overflow errors
package core

import "testing"

func TestWordsToInt(t *testing.T) {
	for _, c := range []int64{0, -1, 0x8000, -0x8001, 1 << 40, -(1 << 40)} {
		if got, err := wordsToInt(intWords(c)); err != nil || got != c {
			t.Fatalf("%d: got %d %v", c, got, err)
		}
	}
	for _, b := range [][]Word{{}, {0, 1}, {0xFFFF, 0x8000}, {1, 0, 0, 0, 0}} {
		if _, err := wordsToInt(b); err == nil {
			t.Fatalf("expected error for %04X", b)
		}
	}
}
