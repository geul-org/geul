//ff:func feature=resolve type=resolver control=sequence
//ff:what dirStore.ReadFile — os.ReadFile of a path under the root
package resolve

import "os"

// ReadFile reads name.
func (s dirStore) ReadFile(name string) ([]byte, error) {
	return os.ReadFile(s.path(name))
}
