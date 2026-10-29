//ff:func feature=resolve type=helper control=sequence
//ff:what mapStore.WriteFile — set a path's bytes
package resolve

import "testing/fstest"

// WriteFile stores b at name.
func (m mapStore) WriteFile(name string, b []byte) error {
	m.MapFS[name] = &fstest.MapFile{Data: append([]byte(nil), b...)}
	return nil
}
