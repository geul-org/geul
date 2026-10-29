//ff:func feature=pack type=helper control=sequence
//ff:what osFS.Rename — os.Rename after making the file world-readable (objects are published files)
package pack

import "os"

// Rename makes from readable (0644) and moves it to to.
func (osFS) Rename(from, to string) error {
	if err := os.Chmod(from, 0o644); err != nil {
		return err
	}
	return os.Rename(from, to)
}
