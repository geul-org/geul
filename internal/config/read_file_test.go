//ff:func feature=config type=helper control=sequence
//ff:what TestReadFile — KEY=VALUE lines with comments and blanks; a missing file is empty; a line without = and an unreadable path fail
package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestReadFile(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "env")
	os.WriteFile(p, []byte("# c\n\nGEULDATA = /d\nGEULFORMAT=json\n"), 0o644)
	m, err := readFile(p)
	if err != nil || m["GEULDATA"] != "/d" || m["GEULFORMAT"] != "json" {
		t.Fatalf("%v %v", m, err)
	}
	if m, err := readFile(filepath.Join(dir, "none")); err != nil || len(m) != 0 {
		t.Fatal("missing")
	}
	os.WriteFile(p, []byte("junk\n"), 0o644)
	if _, err := readFile(p); err == nil {
		t.Fatal("junk: want error")
	}
	if _, err := readFile(dir); err == nil {
		t.Fatal("folder: want error")
	}
}
