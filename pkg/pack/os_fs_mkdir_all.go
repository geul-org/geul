//ff:func feature=pack type=helper control=sequence
//ff:what osFS.MkdirAll — os.MkdirAll with mode 0755
package pack

import "os"

// MkdirAll creates dir and its parents.
func (osFS) MkdirAll(dir string) error {
	return os.MkdirAll(dir, 0o755)
}
