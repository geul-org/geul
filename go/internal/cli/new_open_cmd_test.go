//ff:func feature=cli type=command control=sequence
//ff:what TestNewOpenCmd — a version becomes a draft; a missing document and no data folder fail
package cli

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestNewOpenCmd(t *testing.T) {
	h := newHarness(t)
	addr, _ := h.doc()
	f := filepath.Join(h.data, "o.geul")
	h.must("open", addr+"@1", "-o", f)
	if out := h.must("show", f); strings.Contains(out, "ALIAS") {
		t.Fatal(out)
	}
	if _, err := h.run("", "open", "me/0x0000000000000009", "-o", filepath.Join(h.data, "p.geul")); err == nil {
		t.Fatal("missing: want error")
	}
	h.cwd = t.TempDir()
	if _, err := h.run("", "open", addr, "-o", filepath.Join(h.data, "q.geul")); err == nil {
		t.Fatal("no data folder: want error")
	}
}
