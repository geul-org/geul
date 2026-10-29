//ff:func feature=pack type=writer control=sequence
//ff:what TestArchiveStaged — staged bytes; fs.ErrNotExist otherwise
package pack

import (
	"errors"
	"io/fs"
	"testing"
)

func TestArchiveStaged(t *testing.T) {
	a := &Archive{puts: map[string][]byte{"a": []byte("x")}}
	if b, err := a.staged("a"); err != nil || string(b) != "x" {
		t.Fatal(err)
	}
	if _, err := a.staged("b"); !errors.Is(err, fs.ErrNotExist) {
		t.Fatal(err)
	}
}
