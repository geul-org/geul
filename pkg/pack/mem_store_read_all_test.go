//ff:func feature=pack type=resolver control=sequence
//ff:what memStore.ReadAll — a stored object or NotFound
package pack

import "context"

// ReadAll returns m[name].
func (m memStore) ReadAll(_ context.Context, name string) ([]byte, error) {
	b, ok := m[name]
	if !ok {
		return nil, notFound(name)
	}
	return b, nil
}
