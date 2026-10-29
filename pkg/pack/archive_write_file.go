//ff:func feature=pack type=writer control=sequence
//ff:what Archive.WriteFile — resolve.Store: stage bytes at a path (seen by reads at once, written by Commit)
package pack

import (
	"errors"
	"io/fs"
	"slices"
)

// WriteFile stages b as the bytes of name.
func (a *Archive) WriteFile(name string, b []byte) error {
	if !fs.ValidPath(name) || name == "." {
		return pathErr("write", name, fs.ErrInvalid)
	}
	if a.dir == "" {
		return pathErr("write", name, errors.New("read-only pack archive"))
	}
	a.mu.Lock()
	defer a.mu.Unlock()
	a.puts[name] = append([]byte(nil), b...)
	delete(a.dels, name)
	if i, found := slices.BinarySearch(a.idx, name); a.idx != nil && !found {
		a.idx = slices.Insert(a.idx, i, name)
	}
	return nil
}
