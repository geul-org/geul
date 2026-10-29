//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what Archive.Unpack — write every path back as files under `<root>/.well-known/geul/` (refuses an existing tree); the file count
package pack

import (
	"fmt"
	"os"
	"path/filepath"
)

// Unpack restores the resolve.Dir layout of the whole archive into root.
func (a *Archive) Unpack(root string) (int, error) {
	base := filepath.Join(root, ".well-known", "geul")
	if _, err := os.Stat(base); err == nil {
		return 0, fmt.Errorf("geul/pack: %s already exists", base)
	}
	paths, err := a.Paths()
	if err != nil {
		return 0, err
	}
	for i, p := range paths {
		if err := a.unpackOne(base, p); err != nil {
			return i, err
		}
	}
	return len(paths), nil
}
