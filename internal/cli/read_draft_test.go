//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what TestReadDraft — HEAD and BODY; a missing file, odd bytes and a broken document fail
package cli

import (
	"os"
	"path/filepath"
	"testing"
)

func TestReadDraft(t *testing.T) {
	dir := t.TempDir()
	if d, err := readDraft(brokenDraft(t, dir)); err != nil || len(d.head) != 2 || len(d.body) != 2 {
		t.Fatalf("%+v %v", d, err)
	}
	if _, err := readDraft(filepath.Join(dir, "none")); err == nil {
		t.Fatal("missing: want error")
	}
	for i, b := range [][]byte{{1}, {0, 5}} {
		p := filepath.Join(dir, string(rune('a'+i)))
		os.WriteFile(p, b, 0o644)
		if _, err := readDraft(p); err == nil {
			t.Errorf("%v: want error", b)
		}
	}
}
