//ff:func feature=cli type=command control=iteration dimension=1
//ff:what TestAppHash — 0x and 64 hex digits; a missing document, odd bytes and a broken document fail
package cli

import (
	"os"
	"path/filepath"
	"testing"
)

func TestAppHash(t *testing.T) {
	h := newHarness(t)
	_, draft := h.doc()
	if out := h.must("hash", draft); len(out) != 67 {
		t.Fatal(out)
	}
	if _, err := h.run("", "hash", "me/x"); err == nil {
		t.Fatal("missing: want error")
	}
	for i, b := range [][]byte{{1}, {0, 5}} {
		f := filepath.Join(h.data, string(rune('a'+i))+".geul")
		os.WriteFile(f, b, 0o644)
		if _, err := h.run("", "hash", f); err == nil {
			t.Errorf("%v: want error", b)
		}
	}
}
