//ff:func feature=pack type=helper control=iteration dimension=1
//ff:what dirFile.ReadDir — fs.ReadDirFile: up to n entries (all when n ≤ 0), io.EOF when n > 0 and none are left
package pack

import (
	"io"
	"io/fs"
)

// ReadDir returns the next entries.
func (d *dirFile) ReadDir(n int) ([]fs.DirEntry, error) {
	k := len(d.rest)
	if n > 0 && n < k {
		k = n
	}
	if n > 0 && k == 0 {
		return nil, io.EOF
	}
	out := make([]fs.DirEntry, k)
	for i := range out {
		out[i] = d.rest[i]
	}
	d.rest = d.rest[k:]
	return out, nil
}
