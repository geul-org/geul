//ff:func feature=resolve type=helper control=sequence
//ff:what mapStore.Remove — delete a path, fs.ErrNotExist when absent
package resolve

import "io/fs"

// Remove deletes name.
func (m mapStore) Remove(name string) error {
	if _, ok := m.MapFS[name]; !ok {
		return &fs.PathError{Op: "remove", Path: name, Err: fs.ErrNotExist}
	}
	delete(m.MapFS, name)
	return nil
}
