//ff:func feature=pack type=writer control=sequence
//ff:what TestSplitPaths — Issuer paths and their h/ copies go to the Issuer, other paths home; with no Issuer everything goes to the top set; read errors fail
package pack

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSplitPaths(t *testing.T) {
	root := fixture(t)
	base := filepath.Join(root, ".well-known", "geul")
	os.WriteFile(filepath.Join(base, "h", "0x"+"ff"+"00000000000000000000000000000000000000000000000000000000000000.geul"), []byte{1}, 0o644)
	paths, _ := listTree(root)
	g, err := splitPaths(base, paths)
	if err != nil || len(g) != 1 || len(g["0x0000"]) != 19 {
		t.Fatal(g, err)
	}
	top, _ := splitPaths(base, []string{"issuers.geul"})
	if len(top[""]) != 1 {
		t.Fatal(top)
	}
	if _, err := splitPaths(base, []string{"0x0000/none.geul"}); err == nil {
		t.Fatal("unreadable")
	}
}
