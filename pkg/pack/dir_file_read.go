//ff:func feature=pack type=helper control=sequence
//ff:what dirFile.Read — a directory has no bytes
package pack

import "io/fs"

// Read fails: directories are listed, not read.
func (d *dirFile) Read([]byte) (int, error) {
	return 0, pathErr("read", d.info.name, fs.ErrInvalid)
}
