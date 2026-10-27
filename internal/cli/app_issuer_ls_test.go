//ff:func feature=cli type=command control=sequence
//ff:what TestAppIssuerLs — nothing without keys, one line per key; a broken key file fails
package cli

import (
	"os"
	"path/filepath"
	"testing"
)

func TestAppIssuerLs(t *testing.T) {
	h := newHarness(t)
	if out := h.must("issuer", "ls"); out != "" {
		t.Fatal(out)
	}
	made := h.must("issuer", "new", "--name", "me")
	if out := h.must("issuer", "ls"); out != made {
		t.Fatal(out)
	}
	os.MkdirAll(filepath.Join(h.conf, "geul", "keys"), 0o700)
	os.WriteFile(filepath.Join(h.conf, "geul", "keys", "x.pem"), []byte("x"), 0o600)
	if _, err := h.run("", "issuer", "ls"); err == nil {
		t.Fatal("broken key: want error")
	}
}
