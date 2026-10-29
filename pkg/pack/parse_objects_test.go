//ff:func feature=pack type=decoder control=iteration dimension=1
//ff:what TestParseObjects — entries round-trip; wrong length, unknown kind, reserved bytes and sizes below the header fail
package pack

import (
	"errors"
	"reflect"
	"testing"
)

func TestParseObjects(t *testing.T) {
	objs := []object{{kind: kindData, size: 70, hash: [32]byte{1}}}
	got, err := parseObjects(encodeObjects(objs))
	if err != nil || !reflect.DeepEqual(got, objs) {
		t.Fatal(got, err)
	}
	ok := encodeObjects(objs)
	mut := func(i int, v byte) []byte { b := append([]byte{}, ok...); b[i] = v; return b }
	for i, b := range [][]byte{ok[:10], mut(4, 3), mut(5, 1), mut(12, 10)} {
		if _, err := parseObjects(b); !errors.Is(err, ErrCorrupt) {
			t.Fatal(i, err)
		}
	}
}
