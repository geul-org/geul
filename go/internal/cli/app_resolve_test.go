//ff:func feature=cli type=command control=sequence
//ff:what TestAppResolve — text and JSON; a missing import, a bad GEULPROXY and a missing root fail; without a data folder only GEULPROXY is used
package cli

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestAppResolve(t *testing.T) {
	h := newHarness(t)
	h.env["GEULPROXY"] = "off"
	addr, _ := h.doc()
	f := filepath.Join(h.data, "b.geul")
	h.must("new", "me/0x0000000000000002", "-o", f)
	h.must("meta", f, "IMPORT-DOC", addr+"@1", "0x0001")
	if out := h.must("resolve", f, "--format", "json"); !strings.HasPrefix(out, "[") {
		t.Fatal(out)
	}
	if out := h.must("resolve", f); !strings.Contains(out, "@1 file://") {
		t.Fatal(out)
	}
	h.must("meta", f, "IMPORT-DOC", addr+"@7", "0x0001")
	if _, err := h.run("", "resolve", f); err == nil {
		t.Fatal("missing import: want error")
	}
	h.env["GEULPROXY"] = "ftp://x"
	if _, err := h.run("", "resolve", f); err == nil {
		t.Fatal("bad GEULPROXY: want error")
	}
	if _, err := h.run("", "resolve", "me/0x0000000000000009"); err == nil {
		t.Fatal("missing root: want error")
	}
	h.env["GEULPROXY"] = "off"
	h.cwd = t.TempDir()
	if _, err := h.run("", "resolve", f); err == nil {
		t.Fatal("no data folder and GEULPROXY off: want error")
	}
}
