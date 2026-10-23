//ff:func feature=resolve type=resolver control=sequence
//ff:what Dir.Fetch — read a path from the local tree, ErrNotFound when the file is missing
package resolve

import (
	"context"
	"errors"
	"io/fs"
	"os"
)

// Fetch reads path from the local tree.
func (d Dir) Fetch(_ context.Context, path string) ([]byte, error) {
	b, err := os.ReadFile(d.file(path))
	if errors.Is(err, fs.ErrNotExist) {
		return nil, ErrNotFound
	}
	return b, err
}
