//ff:func feature=cli type=helper control=sequence
//ff:what TestOpenArchive — a packed folder opens; no tree and a directory archive fail
package cli

import (
	"testing"
)

func TestOpenArchive(t *testing.T) {
	h := newHarness(t)
	h.doc()
	if _, err := openArchive(h.data); err == nil {
		t.Fatal("directory archive: want error")
	}
	if _, err := openArchive(t.TempDir()); err == nil {
		t.Fatal("no tree: want error")
	}
	h.packed()
	if a, err := openArchive(h.data); err != nil || len(a.Sets()) != 2 {
		t.Fatal(err)
	}
}
