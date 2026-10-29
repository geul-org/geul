//ff:func feature=pack type=helper control=sequence
//ff:what osFS.CreateTemp — a new `.tmp-*` file in dir
package pack

import "os"

// CreateTemp creates a temp file in dir; renaming it publishes it.
func (osFS) CreateTemp(dir string) (objFile, error) {
	f, err := os.CreateTemp(dir, ".tmp-*")
	if err != nil {
		return nil, err
	}
	return f, nil
}
