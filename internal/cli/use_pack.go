//ff:func feature=cli type=helper control=selection
//ff:what usePack — whether a `.well-known/geul/` directory is (or, being new, becomes) a pack archive
package cli

import (
	"os"

	"github.com/geul-org/gogeul/pkg/pack"
)

// usePack is true for pack sets and for an empty folder (a new archive)
// unless GEULSTORE is dir; a directory tree with files, or no folder at all,
// stays a directory archive.
func usePack(wk, mode string) bool {
	ents, err := os.ReadDir(wk)
	switch {
	case pack.IsPacked(wk):
		return true
	case mode == "dir":
		return false
	default:
		return err == nil && len(ents) == 0
	}
}
