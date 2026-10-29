//ff:func feature=pack type=decoder control=sequence
//ff:what TestWindowSize — exponent and mantissa (0x38 = 128 KiB)
package pack

import "testing"

func TestWindowSize(t *testing.T) {
	if windowSize(0x38) != 128<<10 || windowSize(0x01) != 1024+128 {
		t.Fatal("windowSize")
	}
}
