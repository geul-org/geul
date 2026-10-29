//ff:func feature=pack type=helper control=sequence
//ff:what pathErr — an *fs.PathError for an Archive operation
package pack

import "io/fs"

// pathErr wraps err as fs functions expect.
func pathErr(op, name string, err error) error {
	return &fs.PathError{Op: op, Path: name, Err: err}
}
