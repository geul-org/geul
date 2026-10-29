//ff:func feature=pack type=writer control=sequence
//ff:what listTree — every file path under `<root>/.well-known/geul/`, sorted
package pack

import (
	"path/filepath"
	"sort"
)

// listTree lists an archive directory (resolve.Dir layout).
func listTree(root string) ([]string, error) {
	t := &treeWalk{root: filepath.Join(root, ".well-known", "geul")}
	if err := filepath.WalkDir(t.root, t.visit); err != nil {
		return nil, err
	}
	sort.Strings(t.paths)
	return t.paths, nil
}
