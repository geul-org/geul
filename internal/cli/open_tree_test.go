//ff:func feature=cli type=helper control=sequence
//ff:what TestOpenTree — the directory store, or the pack archive; a broken pack archive fails
package cli

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/geul-org/gogeul/pkg/pack"
)

func TestOpenTree(t *testing.T) {
	h := newHarness(t)
	h.doc()
	if tr, wk, err := openTree(h.data); err != nil || tr == nil || wk != filepath.Join(h.data, ".well-known", "geul") {
		t.Fatal(err)
	}
	h.packed()
	tr, _, err := openTree(h.data)
	if _, ok := tr.(*pack.Archive); err != nil || !ok {
		t.Fatal(err)
	}
	os.WriteFile(filepath.Join(h.data, ".well-known", "geul", "0x0000", "latest.txt"), []byte("x"), 0o644)
	if _, _, err := openTree(h.data); err == nil {
		t.Fatal("broken")
	}
}
