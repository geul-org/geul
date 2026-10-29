//ff:func feature=pack type=writer control=sequence
//ff:what Archive.unpackOne — read one path and write it as a file under base
package pack

import (
	"os"
	"path/filepath"
)

// unpackOne writes path's bytes to base/path.
func (a *Archive) unpackOne(base, path string) error {
	b, err := a.ReadFile(path)
	if err != nil {
		return err
	}
	dst := treeFile(base, path)
	if err := os.MkdirAll(filepath.Dir(dst), 0o755); err != nil {
		return err
	}
	return os.WriteFile(dst, b, 0o644)
}
