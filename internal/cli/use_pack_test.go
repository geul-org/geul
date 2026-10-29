//ff:func feature=cli type=helper control=sequence
//ff:what TestUsePack — pack sets and empty folders are pack (unless GEULSTORE=dir); folders with files and missing folders are directories
package cli

import (
	"os"
	"path/filepath"
	"testing"
)

func TestUsePack(t *testing.T) {
	empty := t.TempDir()
	tree := t.TempDir()
	os.WriteFile(filepath.Join(tree, "issuers.geul"), nil, 0o644)
	packed := t.TempDir()
	os.WriteFile(filepath.Join(packed, "latest.txt"), nil, 0o644)
	if !usePack(empty, "") || usePack(empty, "dir") || usePack(tree, "") || !usePack(packed, "dir") || usePack(filepath.Join(empty, "none"), "") {
		t.Fatal("usePack")
	}
}
