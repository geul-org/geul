//ff:func feature=pack type=helper control=sequence
//ff:what fileInfo.Mode — read-only file or directory mode bits
package pack

import "io/fs"

// Mode is 0444 for files and ModeDir|0555 for directories.
func (f fileInfo) Mode() fs.FileMode {
	if f.dir {
		return fs.ModeDir | 0o555
	}
	return 0o444
}
