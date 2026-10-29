//ff:func feature=resolve type=resolver control=sequence
//ff:what TestDirStoreOpen — files and directories open with fs rules; a missing path is fs.ErrNotExist
package resolve

import (
	"errors"
	"io/fs"
	"testing"
)

func TestDirStoreOpen(t *testing.T) {
	s := dirStore{root: t.TempDir()}
	s.WriteFile("a/b.geul", []byte("x"))
	f, err := s.Open("a/b.geul")
	if err != nil {
		t.Fatal(err)
	}
	f.Close()
	if fi, err := fs.Stat(s, "a"); err != nil || !fi.IsDir() {
		t.Fatal(err)
	}
	if _, err := s.Open("none"); !errors.Is(err, fs.ErrNotExist) {
		t.Fatal(err)
	}
}
