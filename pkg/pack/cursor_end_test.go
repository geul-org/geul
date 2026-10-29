//ff:func feature=pack type=decoder control=sequence
//ff:what TestCursorEnd — nil when consumed cleanly; leftover bytes or a failed read are ErrCorrupt
package pack

import (
	"errors"
	"testing"
)

func TestCursorEnd(t *testing.T) {
	if err := (&cursor{}).end("x"); err != nil {
		t.Fatal(err)
	}
	if err := (&cursor{b: []byte{1}}).end("x"); !errors.Is(err, ErrCorrupt) {
		t.Fatal(err)
	}
	if err := (&cursor{bad: true}).end("x"); !errors.Is(err, ErrCorrupt) {
		t.Fatal(err)
	}
}
