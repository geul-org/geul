//ff:func feature=resolve type=resolver control=sequence
//ff:what Dir.write — store bytes at a path, creating its directories
package resolve

import (
	"os"
	"path/filepath"
)

// write stores b at path in the local tree.
func (d Dir) write(path string, b []byte) error {
	f := d.file(path)
	if err := os.MkdirAll(filepath.Dir(f), 0o755); err != nil {
		return err
	}
	return os.WriteFile(f, b, 0o644)
}
