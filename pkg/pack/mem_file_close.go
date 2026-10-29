//ff:func feature=pack type=helper control=sequence
//ff:what memFile.Close — nothing to release
package pack

// Close does nothing.
func (f *memFile) Close() error { return nil }
