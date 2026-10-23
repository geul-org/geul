//ff:func feature=resolve type=resolver control=sequence
//ff:what Memory.Fetch — the bytes stored at path, ErrNotFound when absent
package resolve

import "context"

// Fetch returns the bytes stored at path.
func (m Memory) Fetch(_ context.Context, path string) ([]byte, error) {
	b, ok := m[path]
	if !ok {
		return nil, ErrNotFound
	}
	return b, nil
}
