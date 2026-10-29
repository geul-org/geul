//ff:func feature=pack type=decoder control=sequence
//ff:what TestReadLE — little-endian of any length up to 8
package pack

import "testing"

func TestReadLE(t *testing.T) {
	if readLE(nil) != 0 || readLE([]byte{1, 2}) != 0x0201 || readLE([]byte{0, 0, 0, 0, 0, 0, 0, 1}) != 1<<56 {
		t.Fatal("readLE")
	}
}
