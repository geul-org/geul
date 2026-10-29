//ff:func feature=pack type=helper control=sequence
//ff:what TestTreeFile — slash paths under base
package pack

import (
	"path/filepath"
	"testing"
)

func TestTreeFile(t *testing.T) {
	if treeFile("b", "h/x.geul") != filepath.Join("b", "h", "x.geul") {
		t.Fatal("treeFile")
	}
}
