//ff:func feature=pack type=helper control=sequence
//ff:what treeFile — the local file of an archive path under a `.well-known/geul/` directory
package pack

import "path/filepath"

// treeFile maps path under base.
func treeFile(base, path string) string {
	return filepath.Join(base, filepath.FromSlash(path))
}
