//ff:func feature=pack type=resolver control=iteration dimension=1
//ff:what IsPacked — whether a `.well-known/geul/` directory holds pack sets (latest.txt at its top or under an Issuer folder)
package pack

import (
	"os"
	"path/filepath"
)

// IsPacked reports whether wk is an Archive rather than a plain directory tree.
func IsPacked(wk string) bool {
	if _, err := os.Stat(filepath.Join(wk, latestFile)); err == nil {
		return true
	}
	ents, _ := os.ReadDir(wk)
	for _, e := range ents {
		if _, err := os.Stat(filepath.Join(wk, e.Name(), latestFile)); err == nil && isIssuer(e.Name()) {
			return true
		}
	}
	return false
}
