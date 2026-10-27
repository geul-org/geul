//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what TestEdgeDefined — assigned codes of each family and reserved gaps
package schema

import (
	"testing"
)

func TestEdgeDefined(t *testing.T) {
	for _, c := range []uint8{0x00, 0x03, 0x06, 0x0B, 0x12, 0x1A, 0x1D, 0x20, 0x59, 0x65, 0x6A, 0x76, 0x79, 0x7F} {
		if !edgeDefined(c) {
			t.Fatalf("0x%02X: want defined", c)
		}
	}
	for _, c := range []uint8{0x07, 0x0C, 0x40, 0x1F, 0x21, 0x26, 0x27, 0x5A, 0x66, 0x77, 0x7A} {
		if edgeDefined(c) {
			t.Fatalf("0x%02X: want reserved", c)
		}
	}
}
