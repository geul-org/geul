//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what TestMetaDefined — assigned codes and reserved gaps
package schema

import (
	"testing"
)

func TestMetaDefined(t *testing.T) {
	for _, c := range []uint8{0x00, 0x0F, 0x14, 0x16, 0x28, 0x39, 0x44, 0x55, 0x60} {
		if !metaDefined(c) {
			t.Fatalf("0x%02X: want defined", c)
		}
	}
	for _, c := range []uint8{0x03, 0x15, 0x17, 0x29, 0x3A, 0x45, 0x56, 0x61, 0x63, 0x64, 0x70} {
		if metaDefined(c) {
			t.Fatalf("0x%02X: want reserved", c)
		}
	}
}
