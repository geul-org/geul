//ff:func feature=pack type=helper control=sequence
//ff:what memFile.Stat — the file's info
package pack

import "io/fs"

// Stat describes the file.
func (f *memFile) Stat() (fs.FileInfo, error) { return f.info, nil }
