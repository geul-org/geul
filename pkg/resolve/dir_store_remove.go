//ff:func feature=resolve type=resolver control=sequence
//ff:what dirStore.Remove — os.Remove of a path under the root
package resolve

import "os"

// Remove deletes name.
func (s dirStore) Remove(name string) error {
	return os.Remove(s.path(name))
}
