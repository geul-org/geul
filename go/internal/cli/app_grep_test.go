//ff:func feature=cli type=command control=sequence
//ff:what TestAppGrep — named documents; a bad pattern, no data folder, a missing document and a broken file fail
package cli

import (
	"os"
	"path/filepath"
	"testing"
)

func TestAppGrep(t *testing.T) {
	h := newHarness(t)
	addr, _ := h.doc()
	if out := h.must("grep", "LABEL", addr+"@1"); out == "" {
		t.Fatal("named document: want a match")
	}
	if _, err := h.run("", "grep", "("); err == nil {
		t.Fatal("bad pattern: want error")
	}
	if _, err := h.run("", "grep", "x", "me/0x0000000000000009"); err == nil {
		t.Fatal("missing: want error")
	}
	bad := filepath.Join(h.data, "bad.geul")
	os.WriteFile(bad, []byte{1}, 0o644)
	if _, err := h.run("", "grep", "x", bad); err == nil {
		t.Fatal("broken: want error")
	}
	h.cwd = t.TempDir()
	if _, err := h.run("", "grep", "x"); err == nil {
		t.Fatal("no data folder: want error")
	}
}
