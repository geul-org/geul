//ff:func feature=server type=helper control=sequence
//ff:what isDir — whether a tree name is a directory
package server

import "io/fs"

// isDir reports whether name in t is a directory.
func isDir(t fs.FS, name string) bool {
	fi, err := fs.Stat(t, name)
	return err == nil && fi.IsDir()
}
