//ff:func feature=server type=helper control=sequence
//ff:what TestIsDir — a directory is, a file and a missing path are not
package server

import (
	"testing"
)

func TestIsDir(t *testing.T) {
	tr := treeOf(treeFixture(t))
	if !isDir(tr, "h") || isDir(tr, "h/a.geul") || isDir(tr, "h/none") {
		t.Fatal("isDir")
	}
}
