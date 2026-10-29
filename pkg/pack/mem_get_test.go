//ff:func feature=pack type=helper control=sequence
//ff:what memGet — a Batch.Get over a map, fs.ErrNotExist for absent paths
package pack

import "io/fs"

// memGet serves bytes from m.
func memGet(m map[string][]byte) func(string) ([]byte, error) {
	return func(p string) ([]byte, error) {
		b, ok := m[p]
		if !ok {
			return nil, fs.ErrNotExist
		}
		return b, nil
	}
}
