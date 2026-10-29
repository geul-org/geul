//ff:func feature=pack type=resolver control=sequence
//ff:what memStore.ReadRange — a slice of a stored object, ErrCorrupt past its end
package pack

import "context"

// ReadRange returns m[name][off:off+n].
func (m memStore) ReadRange(_ context.Context, name string, off, n uint64) ([]byte, error) {
	b, ok := m[name]
	if !ok {
		return nil, notFound(name)
	}
	if off+n > uint64(len(b)) {
		return nil, corrupt("short")
	}
	return b[off : off+n], nil
}
