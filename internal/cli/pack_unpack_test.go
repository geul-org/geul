//ff:func feature=cli type=command control=sequence
//ff:what TestPackUnpack — the pack archive back as the same files; a directory archive and an existing output fail
package cli

import (
	"bytes"
	"path/filepath"
	"strings"
	"testing"
)

func TestPackUnpack(t *testing.T) {
	h := newHarness(t)
	h.doc()
	src := h.data
	h.packed()
	var out bytes.Buffer
	dst := filepath.Join(t.TempDir(), "un")
	if err := packUnpack(h.data, dst, &out); err != nil || !strings.Contains(out.String(), " files → ") {
		t.Fatal(err)
	}
	if a, b := treeFiles(t, src), treeFiles(t, dst); a != b {
		t.Fatalf("differs:\n%s\n%s", a, b)
	}
	if err := packUnpack(src, t.TempDir(), &out); err == nil {
		t.Fatal("directory archive")
	}
	if err := packUnpack(h.data, dst, &out); err == nil {
		t.Fatal("exists")
	}
}
