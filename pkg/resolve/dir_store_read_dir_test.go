//ff:func feature=resolve type=resolver control=sequence
//ff:what TestDirStoreReadDir — entries in name order, "." is the root; missing is fs.ErrNotExist
package resolve

import (
	"errors"
	"io/fs"
	"testing"
)

func TestDirStoreReadDir(t *testing.T) {
	s := dirStore{root: t.TempDir()}
	s.WriteFile("b.geul", nil)
	s.WriteFile("a/x.geul", nil)
	es, err := s.ReadDir(".")
	if err != nil || len(es) != 2 || es[0].Name() != "a" || !es[0].IsDir() {
		t.Fatal(es, err)
	}
	if _, err := s.ReadDir("none"); !errors.Is(err, fs.ErrNotExist) {
		t.Fatal(err)
	}
}
