//ff:func feature=pack type=helper control=sequence
//ff:what dirEntry.Info — the entry as an fs.FileInfo; a listed file is stat'ed through its Archive for its size
package pack

import "io/fs"

// Info describes the entry.
func (e dirEntry) Info() (fs.FileInfo, error) {
	if e.dir || e.a == nil {
		return fileInfo{name: e.name, dir: e.dir, size: e.size}, nil
	}
	return e.a.Stat(e.path)
}
