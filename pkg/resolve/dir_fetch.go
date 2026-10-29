//ff:func feature=resolve type=resolver control=sequence
//ff:what Dir.Fetch — read a path from the archive, ErrNotFound when it is missing
package resolve

import (
	"context"
	"errors"
	"io/fs"
)

// Fetch reads path from the archive's Store.
func (d Dir) Fetch(_ context.Context, path string) ([]byte, error) {
	b, err := d.store().ReadFile(path)
	if errors.Is(err, fs.ErrNotExist) {
		return nil, ErrNotFound
	}
	return b, err
}
