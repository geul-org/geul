//ff:func feature=cli type=helper control=sequence
//ff:what treeFiles — test fixture: every file of root/.well-known/geul with its SHA-256, one per line, sorted
package cli

import (
	"crypto/sha256"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// treeFiles summarizes a directory tree for comparison.
func treeFiles(t *testing.T, root string) string {
	t.Helper()
	base := filepath.Join(root, ".well-known", "geul")
	var lines []string
	filepath.WalkDir(base, func(p string, d fs.DirEntry, err error) error {
		if err == nil && !d.IsDir() {
			b, _ := os.ReadFile(p)
			rel, _ := filepath.Rel(base, p)
			lines = append(lines, fmt.Sprintf("%s %x", rel, sha256.Sum256(b)))
		}
		return err
	})
	return strings.Join(lines, "\n")
}
