//ff:func feature=pack type=decoder control=iteration dimension=1
//ff:what TestParseObjectHeader — short input, magic, version, kind, reserved fields and inventory bounds are checked
package pack

import (
	"errors"
	"testing"
)

func TestParseObjectHeader(t *testing.T) {
	if _, err := parseObjectHeader(encodeObjectHeader(objectHeader{kind: kindIndex, size: 64})); err != nil {
		t.Fatal(err)
	}
	ok := encodeObjectHeader(objectHeader{kind: kindData, size: 100, invOff: 80, invLen: 20})
	mut := func(i int, v byte) []byte { b := append([]byte{}, ok...); b[i] = v; return b }
	bad := [][]byte{ok[:63], mut(0, 'X'), mut(8, 2), mut(10, 3), mut(12, 1), mut(60, 1), mut(32, 21),
		encodeObjectHeader(objectHeader{kind: kindIndex, size: 64, invOff: 64})}
	for i, b := range bad {
		if _, err := parseObjectHeader(b); !errors.Is(err, ErrCorrupt) {
			t.Fatal(i, err)
		}
	}
}
