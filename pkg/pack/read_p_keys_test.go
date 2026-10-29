//ff:func feature=pack type=decoder control=sequence
//ff:what TestReadPKeys — rebuilds keys; short columns and non-increasing keys are ErrCorrupt
package pack

import (
	"errors"
	"testing"
)

func TestReadPKeys(t *testing.T) {
	keys, err := readPKeys(&cursor{b: appendPKeys(nil, []string{pKey(1, 8, 0), pKey(1, 8, 1)})}, 2)
	if err != nil || keys[1] != pKey(1, 8, 1) {
		t.Fatal(keys, err)
	}
	if _, err := readPKeys(&cursor{b: []byte{1}}, 2); !errors.Is(err, ErrCorrupt) {
		t.Fatal(err)
	}
	if _, err := readPKeys(&cursor{b: []byte{1, 0, 8, 0, 0, 0}}, 2); !errors.Is(err, ErrCorrupt) {
		t.Fatal(err)
	}
}
