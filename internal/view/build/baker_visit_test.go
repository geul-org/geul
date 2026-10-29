//ff:func feature=build type=command control=sequence
//ff:what TestBakerVisit — a walk error is passed on; a directory and a document are counted and written; other files are skipped
package build

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"testing"

	"github.com/geul-org/gogeul/pkg/resolve"
)

func TestBakerVisit(t *testing.T) {
	root := fixture(t)
	b := &baker{tree: resolve.Dir{Root: root}.Tree(), src: filepath.Join(root, ".well-known", "geul"), out: filepath.Join(t.TempDir(), ".well-known", "geul")}
	walk := errors.New("walk")
	if err := b.visit(".", nil, walk); err != walk {
		t.Fatalf("walk: %v", err)
	}
	entry := func(p string) fs.DirEntry { fi, _ := os.Stat(p); return fs.FileInfoToDirEntry(fi) }
	if err := b.visit(".", entry(b.src), nil); err != nil {
		t.Fatal(err)
	}
	h := filepath.Join(b.src, "h")
	if err := b.visit("h", entry(h), nil); err != nil {
		t.Fatal(err)
	}
	if err := b.visit("h/a.geul", entry(filepath.Join(h, "a.geul")), nil); err != nil {
		t.Fatal(err)
	}
	if err := b.visit("h/notes.txt", entry(filepath.Join(h, "notes.txt")), nil); err != nil {
		t.Fatal(err)
	}
	if b.count != (Count{Docs: 1, Dirs: 2}) {
		t.Fatalf("count: %+v", b.count)
	}
	if _, err := os.Stat(filepath.Join(b.out, "index.html")); err != nil {
		t.Fatal("top listing missing")
	}
}
