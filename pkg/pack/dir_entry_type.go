//ff:func feature=pack type=helper control=sequence
//ff:what dirEntry.Type — fs.ModeDir for directories, 0 for files
package pack

import "io/fs"

// Type is the entry's type bits.
func (e dirEntry) Type() fs.FileMode {
	if e.dir {
		return fs.ModeDir
	}
	return 0
}
