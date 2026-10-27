//ff:func feature=build type=helper control=sequence
//ff:what TestCheckOut — the source itself and a path inside it are refused; a sibling with the same prefix and an outside path are not
package build

import (
	"path/filepath"
	"testing"
)

func TestCheckOut(t *testing.T) {
	src := filepath.Join(t.TempDir(), "geul")
	if checkOut(src, src) == nil || checkOut(src, filepath.Join(src, "x")) == nil {
		t.Fatal("inside: want error")
	}
	if err := checkOut(src, src+"-out"); err != nil {
		t.Fatal(err)
	}
	if err := checkOut(src, t.TempDir()); err != nil {
		t.Fatal(err)
	}
}
