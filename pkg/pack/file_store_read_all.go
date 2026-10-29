//ff:func feature=pack type=resolver control=sequence
//ff:what FileStore.ReadAll — a whole local object, ErrNotFound when the file is missing
package pack

import (
	"context"
	"os"
)

// ReadAll reads the object name.
func (s FileStore) ReadAll(_ context.Context, name string) ([]byte, error) {
	b, err := os.ReadFile(s.path(name))
	if err != nil {
		return nil, missing(err, name)
	}
	return b, nil
}
