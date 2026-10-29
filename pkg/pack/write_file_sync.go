//ff:func feature=pack type=writer control=sequence
//ff:what writeFileSync — write a file through a synced temp file and a rename
package pack

import (
	"os"
	"path/filepath"
)

// writeFileSync writes b to path atomically: a temp file in the same folder,
// readable by all (0644, as objects are), synced, then renamed over path.
func writeFileSync(path string, b []byte) error {
	f, err := os.CreateTemp(filepath.Dir(path), ".tmp-*")
	if err != nil {
		return err
	}
	_, err = f.Write(b)
	if err == nil {
		err = f.Chmod(0o644)
	}
	if err == nil {
		err = f.Sync()
	}
	if cerr := f.Close(); err == nil {
		err = cerr
	}
	if err == nil {
		err = os.Rename(f.Name(), path)
	}
	if err != nil {
		os.Remove(f.Name())
	}
	return err
}
