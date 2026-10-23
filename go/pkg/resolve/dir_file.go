//ff:func feature=resolve type=resolver control=sequence
//ff:what Dir.file — the local file a `.well-known/geul/` path maps to
package resolve

import "path/filepath"

// file maps path to its file under Root.
func (d Dir) file(path string) string {
	return filepath.Join(d.Root, filepath.FromSlash(wellKnown), filepath.FromSlash(path))
}
