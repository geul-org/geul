//ff:func feature=cli type=helper control=sequence
//ff:what TestCheckPath — a folder's .geul files only; a named file whatever its name; an unreadable file fails
package cli

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCheckPath(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "a.geul"), []byte{0, 5}, 0o644)
	os.WriteFile(filepath.Join(dir, "b.txt"), []byte{0, 5}, 0o644)
	if fs, err := checkPath(dir); err != nil || len(fs) == 0 || fs[0].Path != filepath.Join(dir, "a.geul") {
		t.Fatalf("%+v %v", fs, err)
	}
	if fs, _ := checkPath(filepath.Join(dir, "b.txt")); len(fs) == 0 {
		t.Fatal("named file: want findings")
	}
	locked := filepath.Join(dir, "c.geul")
	os.WriteFile(locked, nil, 0o000)
	if _, err := checkPath(locked); err == nil {
		t.Fatal("unreadable: want error")
	}
}
