//ff:func feature=cli type=command control=sequence
//ff:what TestTreeRoot — the parent of .well-known, the .well-known/geul directory itself, and a directory without a tree
package cli

import (
	"os"
	"path/filepath"
	"testing"
)

func TestTreeRoot(t *testing.T) {
	root := t.TempDir()
	geul := filepath.Join(root, ".well-known", "geul")
	os.MkdirAll(geul, 0o755)
	if r, err := treeRoot(root); err != nil || r != root {
		t.Fatalf("parent: %q %v", r, err)
	}
	if r, err := treeRoot(geul); err != nil || r != root {
		t.Fatalf("tree: %q %v", r, err)
	}
	if _, err := treeRoot(t.TempDir()); err == nil {
		t.Fatal("none: want error")
	}
}
