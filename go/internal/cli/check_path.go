//ff:func feature=cli type=helper control=sequence
//ff:what checkPath — schema findings of one file, or of every .geul file under a folder
package cli

import (
	"io/fs"
	"path/filepath"
	"strings"
)

// checkPath walks p (a file or a folder) and checks each `.geul` file
// (checkFile).
func checkPath(p string) ([]finding, error) {
	var out []finding
	err := filepath.WalkDir(p, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() || (path != p && !strings.HasSuffix(path, ".geul")) {
			return err
		}
		got, err := checkFile(path)
		out = append(out, got...)
		return err
	})
	return out, err
}
