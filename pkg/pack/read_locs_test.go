//ff:func feature=pack type=decoder control=sequence
//ff:what TestReadLocs — round trip; short columns and invalid entries are ErrCorrupt
package pack

import (
	"errors"
	"reflect"
	"testing"
)

func TestReadLocs(t *testing.T) {
	ls := []Loc{{Obj: 1, Off: 64, Len: 4}, {Obj: 1, Off: 68, Len: 2}}
	got, err := readLocs(&cursor{b: appendLocs(nil, ls)}, 2, true)
	if err != nil || !reflect.DeepEqual(got, ls) {
		t.Fatal(got, err)
	}
	if _, err := readLocs(&cursor{b: []byte{1}}, 1, true); !errors.Is(err, ErrCorrupt) {
		t.Fatal(err)
	}
	if _, err := readLocs(&cursor{b: []byte{0, 0, 0, 0}}, 1, true); !errors.Is(err, ErrCorrupt) {
		t.Fatal(err)
	}
}
