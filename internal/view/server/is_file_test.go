//ff:func feature=server type=helper control=sequence
//ff:what TestIsFile — a regular file is, a directory and a missing path are not
package server

import (
	"testing"
)

func TestIsFile(t *testing.T) {
	tr := treeOf(treeFixture(t))
	if !isFile(tr, "h/a.geul") || isFile(tr, "h") || isFile(tr, "h/none") {
		t.Fatal("isFile")
	}
}
