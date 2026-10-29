//ff:func feature=pack type=decoder control=sequence
//ff:what TestParseDicts — round trip; a length that is not 52 per entry fails
package pack

import (
	"errors"
	"reflect"
	"testing"
)

func TestParseDicts(t *testing.T) {
	ds := []dictEntry{{id: 40000, loc: Loc{Obj: 1, Off: 64, Len: 9}, hash: [32]byte{5}}}
	got, err := parseDicts(encodeDicts(ds))
	if err != nil || !reflect.DeepEqual(got, ds) {
		t.Fatal(got, err)
	}
	if _, err := parseDicts([]byte{1, 0, 0, 0}); !errors.Is(err, ErrCorrupt) {
		t.Fatal(err)
	}
}
