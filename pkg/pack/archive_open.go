//ff:func feature=pack type=resolver control=sequence
//ff:what Archive.Open — fs.FS: a directory as a ReadDirFile, a file as a seekable in-memory file
package pack

import (
	"bytes"
	"io/fs"
	"path"
)

// Open opens name ("." is the top).
func (a *Archive) Open(name string) (fs.File, error) {
	if !fs.ValidPath(name) {
		return nil, pathErr("open", name, fs.ErrInvalid)
	}
	if a.isDir(name) {
		a.mu.Lock()
		idx, _ := a.index()
		a.mu.Unlock()
		return &dirFile{info: fileInfo{name: path.Base(name), dir: true}, rest: a.entries(idx, name)}, nil
	}
	b, err := a.ReadFile(name)
	if err != nil {
		return nil, err
	}
	return &memFile{info: fileInfo{name: path.Base(name), size: int64(len(b))}, Reader: bytes.NewReader(b)}, nil
}
