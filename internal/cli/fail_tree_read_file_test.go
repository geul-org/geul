//ff:func feature=cli type=helper control=sequence
//ff:what failTree.ReadFile — always fails
package cli

import "errors"

// ReadFile fails.
func (failTree) ReadFile(string) ([]byte, error) {
	return nil, errors.New("unreadable")
}
