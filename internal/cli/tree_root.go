//ff:func feature=cli type=command control=selection
//ff:what treeRoot — the root of a .well-known/geul tree from a directory argument
package cli

import (
	"fmt"
	"os"
	"path/filepath"
)

// treeRoot accepts either the directory that holds .well-known/geul or that
// .well-known/geul directory itself, and returns the former.
func treeRoot(dir string) (string, error) {
	_, err := os.Stat(filepath.Join(dir, ".well-known", "geul"))
	switch {
	case err == nil:
		return dir, nil
	case filepath.Base(dir) == "geul" && filepath.Base(filepath.Dir(dir)) == ".well-known":
		return filepath.Dir(filepath.Dir(dir)), nil
	default:
		return "", fmt.Errorf("%s: no .well-known/geul tree", dir)
	}
}
