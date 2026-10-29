//ff:func feature=cli type=helper control=sequence
//ff:what failTree.WriteFile — always fails
package cli

import "errors"

// WriteFile fails.
func (failTree) WriteFile(string, []byte) error {
	return errors.New("read-only")
}
