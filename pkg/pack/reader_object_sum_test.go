//ff:func feature=pack type=rule control=sequence
//ff:what TestReaderObjectSum — SHA-256 of an object read in chunks; a short object fails
package pack

import (
	"context"
	"errors"
	"testing"
)

func TestReaderObjectSum(t *testing.T) {
	obj, o := objBytes(kindIndex, []byte("abc"), nil)
	r := newReader(memStore{objectName(o): obj}, Options{}, &manifest{}, "")
	if sum, err := r.objectSum(context.Background(), o); err != nil || sum != o.hash {
		t.Fatal(err)
	}
	o.size++
	if _, err := r.objectSum(context.Background(), o); !errors.Is(err, ErrCorrupt) {
		t.Fatal(err)
	}
}
