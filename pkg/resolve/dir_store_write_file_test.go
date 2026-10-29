//ff:func feature=resolve type=resolver control=sequence
//ff:what TestDirStoreWriteFile — parents are created; a file in the way fails
package resolve

import (
	"testing"
)

func TestDirStoreWriteFile(t *testing.T) {
	s := dirStore{root: t.TempDir()}
	if err := s.WriteFile("a/b/c.geul", []byte("x")); err != nil {
		t.Fatal(err)
	}
	if b, _ := s.ReadFile("a/b/c.geul"); string(b) != "x" {
		t.Fatal(string(b))
	}
	if err := s.WriteFile("a/b/c.geul/d", nil); err == nil {
		t.Fatal("file in the way")
	}
}
