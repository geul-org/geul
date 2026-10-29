//ff:func feature=cli type=helper control=sequence
//ff:what checkTree — schema findings of every .geul path of an archive tree (directory or pack sets), named under base
package cli

import (
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/geul-org/gogeul/pkg/resolve"
)

// checkTree walks t in path order and checks each `.geul` (checkBytes); a
// finding's path is the file the path would be in a directory tree at base.
func checkTree(t resolve.Store, base string) ([]finding, error) {
	var out []finding
	err := fs.WalkDir(t, ".", func(p string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() || !strings.HasSuffix(p, ".geul") {
			return err
		}
		b, err := t.ReadFile(p)
		out = append(out, checkBytes(filepath.Join(base, filepath.FromSlash(p)), b)...)
		return err
	})
	return out, err
}
