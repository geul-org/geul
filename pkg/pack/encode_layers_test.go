//ff:func feature=pack type=encoder control=sequence
//ff:what TestEncodeLayers — count, encoding, reserved, root length and bytes per layer
package pack

import "testing"

func TestEncodeLayers(t *testing.T) {
	b := encodeLayers([]layer{{count: 5, enc: 1, root: []byte("xy")}})
	if len(b) != 4+8+4+4+2 || b[4] != 5 || b[12] != 1 || b[16] != 2 || string(b[20:]) != "xy" {
		t.Fatalf("%x", b)
	}
}
