//ff:func feature=core type=decoder control=iteration dimension=1
//ff:what TestWordsToUint — value read plus empty, padded and overflow errors
package core

import "testing"

func TestWordsToUint(t *testing.T) {
	if n, err := wordsToUint([]Word{1, 0}); err != nil || n != 0x10000 {
		t.Fatalf("got %d %v", n, err)
	}
	for _, b := range [][]Word{{}, {0, 1}, {1, 0, 0, 0, 0}} {
		if _, err := wordsToUint(b); err == nil {
			t.Fatalf("expected error for %v", b)
		}
	}
}
