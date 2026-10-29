//ff:func feature=pack type=helper control=sequence
//ff:what dirFile.Close — nothing to release
package pack

// Close does nothing.
func (d *dirFile) Close() error { return nil }
