//ff:func feature=server type=helper control=sequence
//ff:what isFile — whether a tree name is a regular file
package server

import "io/fs"

// isFile reports whether name in t is a regular file.
func isFile(t fs.FS, name string) bool {
	fi, err := fs.Stat(t, name)
	return err == nil && fi.Mode().IsRegular()
}
