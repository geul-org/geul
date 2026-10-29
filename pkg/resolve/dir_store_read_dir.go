//ff:func feature=resolve type=resolver control=sequence
//ff:what dirStore.ReadDir — os.ReadDir of a directory under the root
package resolve

import (
	"io/fs"
	"os"
)

// ReadDir lists name in name order.
func (s dirStore) ReadDir(name string) ([]fs.DirEntry, error) {
	return os.ReadDir(s.path(name))
}
