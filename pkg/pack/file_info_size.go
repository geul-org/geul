//ff:func feature=pack type=helper control=sequence
//ff:what fileInfo.Size — the byte length (0 for directories)
package pack

// Size is the file length.
func (f fileInfo) Size() int64 { return f.size }
