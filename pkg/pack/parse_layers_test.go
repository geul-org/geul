//ff:func feature=pack type=decoder control=iteration dimension=1
//ff:what TestParseLayers — round trip; bad count, bad encoding, reserved bytes, truncation and trailing bytes fail
package pack

import (
	"errors"
	"reflect"
	"testing"
)

func TestParseLayers(t *testing.T) {
	ls := []layer{{count: 2, enc: 1, root: []byte("r")}}
	got, err := parseLayers(encodeLayers(ls))
	if err != nil || !reflect.DeepEqual(got, ls) {
		t.Fatal(got, err)
	}
	ok := encodeLayers(ls)
	mut := func(i int, v byte) []byte { b := append([]byte{}, ok...); b[i] = v; return b }
	bad := [][]byte{{9, 0, 0, 0}, mut(12, 2), mut(13, 1), ok[:len(ok)-1], append(append([]byte{}, ok...), 0)}
	for i, b := range bad {
		if _, err := parseLayers(b); !errors.Is(err, ErrCorrupt) {
			t.Fatal(i, err)
		}
	}
}
