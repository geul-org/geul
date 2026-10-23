//ff:func feature=resolve type=resolver control=sequence
//ff:what Dir.file — the local file a `.well-known/geul/` path maps to
package resolve

import (
	"path/filepath"
	"strings"
)

// file maps path to its file under Root; a version list path gets the list file.
func (d Dir) file(path string) string {
	if strings.HasSuffix(path, "/") {
		path += listFile
	}
	return filepath.Join(d.Root, filepath.FromSlash(wellKnown), filepath.FromSlash(path))
}
