//ff:func feature=server type=helper control=sequence
//ff:what TestIsFile — a regular file is, a directory and a missing path are not
package server

import (
	"path/filepath"
	"testing"
)

func TestIsFile(t *testing.T) {
	h := filepath.Join(treeFixture(t), ".well-known", "geul", "h")
	if !isFile(filepath.Join(h, "a.geul")) || isFile(h) || isFile(filepath.Join(h, "none")) {
		t.Fatal("isFile")
	}
}
