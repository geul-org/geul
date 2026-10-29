//ff:type feature=pack type=model
//ff:what dirFile — an opened Archive directory: its info and the entries not yet returned by ReadDir
package pack

// dirFile is the fs.ReadDirFile Open returns for a directory.
type dirFile struct {
	info fileInfo
	rest []dirEntry
}
