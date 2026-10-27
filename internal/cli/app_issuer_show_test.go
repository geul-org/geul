//ff:func feature=cli type=command control=sequence
//ff:what TestAppIssuerShow — the matching key among several; an unknown Issuer and a broken key file fail
package cli

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestAppIssuerShow(t *testing.T) {
	h := newHarness(t)
	if _, err := h.run("", "issuer", "show", "me"); err == nil {
		t.Fatal("unknown: want error")
	}
	h.must("issuer", "new", "--name", "a")
	h.must("issuer", "new", "--name", "me")
	if out := h.must("issuer", "show", "me"); !strings.HasPrefix(out, "name me\n") {
		t.Fatal(out)
	}
	os.MkdirAll(filepath.Join(h.conf, "geul", "keys"), 0o700)
	os.WriteFile(filepath.Join(h.conf, "geul", "keys", "x.pem"), []byte("x"), 0o600)
	if _, err := h.run("", "issuer", "show", "x"); err == nil {
		t.Fatal("broken key: want error")
	}
}
