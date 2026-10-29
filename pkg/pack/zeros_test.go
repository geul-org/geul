//ff:func feature=pack type=rule control=sequence
//ff:what TestZeros — reserved bytes must all be zero
package pack

import "testing"

func TestZeros(t *testing.T) {
	if !zeros(nil) || !zeros([]byte{0, 0}) || zeros([]byte{0, 1}) {
		t.Fatal("zeros")
	}
}
