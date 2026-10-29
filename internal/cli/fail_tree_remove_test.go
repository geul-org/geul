//ff:func feature=cli type=helper control=sequence
//ff:what failTree.Remove — always fails
package cli

import "errors"

// Remove fails.
func (failTree) Remove(string) error {
	return errors.New("read-only")
}
