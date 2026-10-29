//ff:func feature=pack type=writer control=sequence
//ff:what Archive.Remove — resolve.Store: stage the removal of a path; fs.ErrNotExist when it is not there
package pack

import (
	"errors"
	"slices"
)

// Remove stages the deletion of name.
func (a *Archive) Remove(name string) error {
	if a.dir == "" {
		return pathErr("remove", name, errors.New("read-only pack archive"))
	}
	if _, err := a.ReadFile(name); err != nil {
		return err
	}
	a.mu.Lock()
	defer a.mu.Unlock()
	delete(a.puts, name)
	a.dels[name] = true
	if i, found := slices.BinarySearch(a.idx, name); found {
		a.idx = slices.Delete(a.idx, i, i+1)
	}
	return nil
}
