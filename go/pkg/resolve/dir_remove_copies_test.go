//ff:func feature=resolve type=resolver control=sequence
//ff:what TestDirRemoveCopies — listed copies go, kept and missing ones are fine; a copy that cannot be removed fails
package resolve

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDirRemoveCopies(t *testing.T) {
	d := Dir{Root: t.TempDir()}
	a, b, c := [32]byte{1}, [32]byte{2}, [32]byte{3}
	os.MkdirAll(filepath.Dir(d.file(hashPath(a))), 0o755)
	os.WriteFile(d.file(hashPath(a)), nil, 0o644)
	os.WriteFile(d.file(hashPath(b)), nil, 0o644)
	if err := d.removeCopies([][32]byte{a, b, c}, map[[32]byte]bool{b: true}); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(d.file(hashPath(a))); !os.IsNotExist(err) {
		t.Fatal("a left")
	}
	if _, err := os.Stat(d.file(hashPath(b))); err != nil {
		t.Fatal("kept b removed")
	}
	os.Mkdir(d.file(hashPath(c)), 0o755)
	os.WriteFile(filepath.Join(d.file(hashPath(c)), "x"), nil, 0o644)
	if err := d.removeCopies([][32]byte{c}, nil); err == nil {
		t.Fatal("non-empty folder: want error")
	}
}
