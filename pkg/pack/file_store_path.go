//ff:func feature=pack type=helper control=sequence
//ff:what FileStore.path — the local file of an object name
package pack

import "path/filepath"

// path maps a slash-separated object name to its file under Root.
func (s FileStore) path(name string) string {
	return filepath.Join(s.Root, filepath.FromSlash(name))
}
