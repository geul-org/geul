//ff:func feature=resolve type=resolver control=sequence
//ff:what TestDirStoreRemove — removes a file; missing is fs.ErrNotExist
package resolve

import (
	"errors"
	"io/fs"
	"testing"
)

func TestDirStoreRemove(t *testing.T) {
	s := dirStore{root: t.TempDir()}
	s.WriteFile("a.geul", nil)
	if err := s.Remove("a.geul"); err != nil {
		t.Fatal(err)
	}
	if err := s.Remove("a.geul"); !errors.Is(err, fs.ErrNotExist) {
		t.Fatal(err)
	}
}
