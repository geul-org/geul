//ff:func feature=config type=helper control=iteration dimension=1
//ff:what Discover — the nearest folder at or above dir that holds .well-known/geul/
package config

import (
	"os"
	"path/filepath"
)

// Discover walks from dir up to the file system root and returns the first
// folder with a `.well-known/geul` directory (as git finds `.git`); "" when
// there is none.
func Discover(dir string) string {
	for d := filepath.Clean(dir); ; d = filepath.Dir(d) {
		if fi, err := os.Stat(filepath.Join(d, ".well-known", "geul")); err == nil && fi.IsDir() {
			return d
		}
		if filepath.Dir(d) == d {
			return ""
		}
	}
}
