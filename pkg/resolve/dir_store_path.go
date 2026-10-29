//ff:func feature=resolve type=helper control=sequence
//ff:what dirStore.path — the local file of a slash path
package resolve

import "path/filepath"

// path maps name under root.
func (s dirStore) path(name string) string {
	return filepath.Join(s.root, filepath.FromSlash(name))
}
