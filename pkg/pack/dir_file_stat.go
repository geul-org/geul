//ff:func feature=pack type=helper control=sequence
//ff:what dirFile.Stat — the directory's info
package pack

import "io/fs"

// Stat describes the directory.
func (d *dirFile) Stat() (fs.FileInfo, error) { return d.info, nil }
