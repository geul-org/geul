//ff:func feature=resolve type=resolver control=sequence
//ff:what TestDirStoreReadFile — bytes of a file; fs.ErrNotExist when missing
package resolve

import (
	"errors"
	"io/fs"
	"testing"
)

func TestDirStoreReadFile(t *testing.T) {
	s := dirStore{root: t.TempDir()}
	s.WriteFile("a.geul", []byte("x"))
	if b, err := s.ReadFile("a.geul"); err != nil || string(b) != "x" {
		t.Fatal(err)
	}
	if _, err := s.ReadFile("b.geul"); !errors.Is(err, fs.ErrNotExist) {
		t.Fatal(err)
	}
}
