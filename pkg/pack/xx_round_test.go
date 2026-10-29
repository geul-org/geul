//ff:func feature=pack type=helper control=sequence
//ff:what TestXxRound — the XXH64 round of zero is zero; nonzero input changes the lane
package pack

import "testing"

func TestXxRound(t *testing.T) {
	if xxRound(0, 0) != 0 || xxRound(0, 1) == 0 {
		t.Fatal("round")
	}
}
