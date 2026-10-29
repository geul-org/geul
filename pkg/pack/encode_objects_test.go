//ff:func feature=pack type=encoder control=sequence
//ff:what TestEncodeObjects — u32 count then 48-byte entries
package pack

import "testing"

func TestEncodeObjects(t *testing.T) {
	b := encodeObjects([]object{{kind: kindIndex, size: 64, hash: [32]byte{9}}})
	if len(b) != 4+48 || b[0] != 1 || b[4] != kindIndex || b[12] != 64 || b[20] != 9 {
		t.Fatalf("%x", b)
	}
}
