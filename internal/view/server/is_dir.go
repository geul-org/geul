//ff:func feature=server type=helper control=sequence
//ff:what isDir — whether a path is a directory
package server

import "os"

// isDir reports whether p names a directory.
func isDir(p string) bool {
	fi, err := os.Stat(p)
	return err == nil && fi.IsDir()
}
