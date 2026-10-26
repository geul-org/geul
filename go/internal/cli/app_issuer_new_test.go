//ff:func feature=cli type=command control=iteration dimension=1
//ff:what TestAppIssuerNew — the default name, a key file only the user reads; a bad name, an existing name and an unwritable folder fail
package cli

import (
	"os"
	"path/filepath"
	"testing"
)

func TestAppIssuerNew(t *testing.T) {
	h := newHarness(t)
	h.must("issuer", "new")
	fi, err := os.Stat(filepath.Join(h.conf, "geul", "keys", "default.pem"))
	if err != nil || fi.Mode().Perm() != 0o600 {
		t.Fatalf("%v %v", fi, err)
	}
	for _, n := range []string{"Bad", "default"} {
		if _, err := h.run("", "issuer", "new", "--name", n); err == nil {
			t.Errorf("%s: want error", n)
		}
	}
	os.Chmod(filepath.Join(h.conf, "geul", "keys"), 0o500)
	defer os.Chmod(filepath.Join(h.conf, "geul", "keys"), 0o700)
	if _, err := h.run("", "issuer", "new", "--name", "other"); err == nil {
		t.Fatal("unwritable: want error")
	}
	h2 := newHarness(t)
	os.MkdirAll(filepath.Join(h2.conf, "geul"), 0o755)
	os.WriteFile(filepath.Join(h2.conf, "geul", "keys"), nil, 0o644)
	if _, err := h2.run("", "issuer", "new", "--name", "x"); err == nil {
		t.Fatal("keys folder is a file: want error")
	}
}
