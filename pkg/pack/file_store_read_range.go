//ff:func feature=pack type=resolver control=sequence
//ff:what FileStore.ReadRange — n bytes at off by pread; a short read (truncated object) is ErrCorrupt
package pack

import (
	"context"
	"fmt"
	"os"
)

// ReadRange reads exactly n bytes at off.
func (s FileStore) ReadRange(_ context.Context, name string, off, n uint64) ([]byte, error) {
	if n == 0 {
		return []byte{}, nil
	}
	f, err := os.Open(s.path(name))
	if err != nil {
		return nil, missing(err, name)
	}
	defer f.Close()
	b := make([]byte, n)
	m, err := f.ReadAt(b, int64(off))
	if uint64(m) == n {
		return b, nil
	}
	return nil, fmt.Errorf("%w: %s: read %d of %d bytes at %d: %v", ErrCorrupt, name, m, n, off, err)
}
