//ff:func feature=pack type=decoder control=sequence
//ff:what TestFcsSize — flag 0 is 1 byte only in single-segment frames, else 0/2/4/8
package pack

import "testing"

func TestFcsSize(t *testing.T) {
	if fcsSize(0, true) != 1 || fcsSize(0, false) != 0 || fcsSize(1, false) != 2 || fcsSize(2, true) != 4 || fcsSize(3, false) != 8 {
		t.Fatal("fcsSize")
	}
}
