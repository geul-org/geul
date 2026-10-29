//ff:func feature=resolve type=resolver control=sequence
//ff:what dirStore.Open — fs.FS Open of a file or directory under the root
package resolve

import (
	"io/fs"
	"os"
)

// Open opens name with os.DirFS rules.
func (s dirStore) Open(name string) (fs.File, error) {
	return os.DirFS(s.root).Open(name)
}
