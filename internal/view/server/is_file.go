//ff:func feature=server type=helper control=sequence
//ff:what isFile — whether a path is a regular file
package server

import "os"

// isFile reports whether p names a regular file.
func isFile(p string) bool {
	fi, err := os.Stat(p)
	return err == nil && fi.Mode().IsRegular()
}
