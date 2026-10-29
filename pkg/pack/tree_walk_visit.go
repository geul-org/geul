//ff:func feature=pack type=writer control=sequence
//ff:what treeWalk.visit — fs.WalkDirFunc that keeps every non-directory as a relative slash path
package pack

import (
	"io/fs"
	"path/filepath"
)

// visit records p when it is not a directory.
func (t *treeWalk) visit(p string, d fs.DirEntry, err error) error {
	if err != nil || d.IsDir() {
		return err
	}
	rel, err := filepath.Rel(t.root, p)
	t.paths = append(t.paths, filepath.ToSlash(rel))
	return err
}
