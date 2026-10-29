//ff:func feature=pack type=resolver control=iteration dimension=1
//ff:what Archive.ReadDir — fs.ReadDirFS: a directory's entries in name order; not a directory is fs.ErrNotExist
package pack

import "io/fs"

// ReadDir lists name. "." is the top of `.well-known/geul/`.
func (a *Archive) ReadDir(name string) ([]fs.DirEntry, error) {
	if !fs.ValidPath(name) {
		return nil, pathErr("readdir", name, fs.ErrInvalid)
	}
	a.mu.Lock()
	idx, err := a.index()
	a.mu.Unlock()
	if err != nil {
		return nil, err
	}
	es := a.entries(idx, name)
	if len(es) == 0 && name != "." {
		return nil, pathErr("readdir", name, fs.ErrNotExist)
	}
	out := make([]fs.DirEntry, len(es))
	for i := range es {
		out[i] = es[i]
	}
	return out, nil
}
