//ff:func feature=pack type=helper control=sequence
//ff:what osFS.Remove — os.Remove
package pack

import "os"

// Remove deletes name.
func (osFS) Remove(name string) error {
	return os.Remove(name)
}
