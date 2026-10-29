//ff:func feature=pack type=encoder control=sequence
//ff:what TestEncodeNamed — length-prefixed path, encoding, reserved, object, offset, length
package pack

import "testing"

func TestEncodeNamed(t *testing.T) {
	b := encodeNamed([]namedEntry{{"ab", Loc{Obj: 1, Off: 64, Len: 3, Enc: 1}}})
	if len(b) != 4+2+2+2+4+8+4 || b[4] != 2 || string(b[6:8]) != "ab" || b[8] != 1 || b[10] != 1 {
		t.Fatalf("%x", b)
	}
}
