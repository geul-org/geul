//ff:func feature=cli type=command control=sequence
//ff:what TestAppDef — the node ID is printed; a missing draft, an object that does not read and a draft breaking the spec fail
package cli

import (
	"path/filepath"
	"testing"
)

func TestAppDef(t *testing.T) {
	h := newHarness(t)
	_, draft := h.doc()
	if out := h.must("def", draft, "Q5"); out != "#5\n" {
		t.Fatal(out)
	}
	if _, err := h.run("", "def", filepath.Join(h.data, "none"), "Q1"); err == nil {
		t.Fatal("missing draft: want error")
	}
	if _, err := h.run("", "def", draft, "word"); err == nil {
		t.Fatal("bad object: want error")
	}
	if _, err := h.run("", "def", brokenDraft(t, h.data), "Q1"); err == nil {
		t.Fatal("draft breaking the spec: want error")
	}
}
