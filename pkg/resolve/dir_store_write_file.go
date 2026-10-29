//ff:func feature=resolve type=resolver control=sequence
//ff:what dirStore.WriteFile — store bytes at a path, creating its directories
package resolve

import (
	"os"
	"path/filepath"
)

// WriteFile writes b to name.
func (s dirStore) WriteFile(name string, b []byte) error {
	f := s.path(name)
	if err := os.MkdirAll(filepath.Dir(f), 0o755); err != nil {
		return err
	}
	return os.WriteFile(f, b, 0o644)
}
