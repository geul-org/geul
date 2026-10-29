//ff:type feature=pack type=model
//ff:what fileInfo — an fs.FileInfo of an Archive path (read-only, no modification time)
package pack

// fileInfo describes a file or directory.
type fileInfo struct {
	name string
	dir  bool
	size int64
}
