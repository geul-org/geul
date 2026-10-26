//ff:func feature=config type=helper control=iteration dimension=1
//ff:what TestWrite — set adds or replaces one key keeping others; empty removes it; unknown keys, GEULENV and unwritable paths fail
package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWrite(t *testing.T) {
	p := filepath.Join(t.TempDir(), "geul", "env")
	if err := Write(p, "GEULDATA", "/a"); err != nil {
		t.Fatal(err)
	}
	Write(p, "GEULISSUER", "me")
	Write(p, "GEULDATA", "/b")
	m, _ := readFile(p)
	if m["GEULDATA"] != "/b" || m["GEULISSUER"] != "me" || len(m) != 2 {
		t.Fatalf("%v", m)
	}
	Write(p, "GEULISSUER", "")
	if m, _ := readFile(p); len(m) != 1 {
		t.Fatalf("unset: %v", m)
	}
	for _, k := range []string{"NOPE", "GEULENV"} {
		if err := Write(p, k, "x"); err == nil {
			t.Errorf("%s: want error", k)
		}
	}
	if err := Write(filepath.Dir(p), "GEULDATA", "x"); err == nil {
		t.Fatal("folder as file: want error")
	}
	blocked := filepath.Join(t.TempDir(), "f")
	os.WriteFile(blocked, nil, 0o644)
	if err := Write(filepath.Join(blocked, "env"), "GEULDATA", "x"); err == nil {
		t.Fatal("parent is a file: want error")
	}
	ro := t.TempDir()
	os.Chmod(ro, 0o500)
	defer os.Chmod(ro, 0o755)
	if err := Write(filepath.Join(ro, "sub", "env"), "GEULDATA", "x"); err == nil {
		t.Fatal("read-only parent: want error")
	}
}
