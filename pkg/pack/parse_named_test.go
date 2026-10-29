//ff:func feature=pack type=decoder control=iteration dimension=1
//ff:what TestParseNamed — round trip; bad count, unsorted or duplicate paths, bad encoding, reserved byte and trailing bytes fail
package pack

import (
	"errors"
	"reflect"
	"testing"
)

func TestParseNamed(t *testing.T) {
	ns := []namedEntry{{"a", Loc{Off: 64, Len: 1}}, {"b", Loc{Off: 65, Len: 1, Enc: 1}}}
	got, err := parseNamed(encodeNamed(ns))
	if err != nil || !reflect.DeepEqual(got, ns) {
		t.Fatal(got, err)
	}
	ok := encodeNamed(ns)
	mut := func(i int, v byte) []byte { b := append([]byte{}, ok...); b[i] = v; return b }
	bad := [][]byte{{9, 0, 0, 0}, encodeNamed([]namedEntry{ns[1], ns[0]}), mut(7, 2), mut(8, 1), append(append([]byte{}, ok...), 0), ok[:30]}
	for i, b := range bad {
		if _, err := parseNamed(b); !errors.Is(err, ErrCorrupt) {
			t.Fatal(i, err)
		}
	}
}
