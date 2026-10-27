//ff:func feature=server type=helper control=sequence
//ff:what TestIsDir — a directory is, a file and a missing path are not
package server

import (
	"path/filepath"
	"testing"
)

func TestIsDir(t *testing.T) {
	h := filepath.Join(treeFixture(t), ".well-known", "geul", "h")
	if !isDir(h) || isDir(filepath.Join(h, "a.geul")) || isDir(filepath.Join(h, "none")) {
		t.Fatal("isDir")
	}
}
