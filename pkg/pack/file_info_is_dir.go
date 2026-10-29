//ff:func feature=pack type=helper control=sequence
//ff:what fileInfo.IsDir — whether it is a directory
package pack

// IsDir reports a directory.
func (f fileInfo) IsDir() bool { return f.dir }
