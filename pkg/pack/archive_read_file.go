//ff:func feature=pack type=resolver control=sequence
//ff:what Archive.ReadFile — fs.ReadFileFS: staged bytes, staged removal, else the sets; a directory or nothing is an error
package pack

import (
	"errors"
	"io/fs"
)

// ReadFile returns the bytes of name (fs.ErrNotExist when absent).
func (a *Archive) ReadFile(name string) ([]byte, error) {
	if !fs.ValidPath(name) {
		return nil, pathErr("read", name, fs.ErrInvalid)
	}
	a.mu.Lock()
	b, staged := a.puts[name]
	gone := a.dels[name]
	a.mu.Unlock()
	if staged {
		return append([]byte(nil), b...), nil
	}
	if gone {
		return nil, pathErr("read", name, fs.ErrNotExist)
	}
	b, ok, err := a.fetch(name)
	if err != nil || ok {
		return b, err
	}
	if a.isDir(name) {
		return nil, pathErr("read", name, errors.New("is a directory"))
	}
	return nil, pathErr("read", name, fs.ErrNotExist)
}
