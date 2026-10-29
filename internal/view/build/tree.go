//ff:func feature=build type=builder control=sequence
//ff:what Tree — bake a directory archive (root/.well-known/geul) into a static copy under dst
package build

import (
	"path/filepath"

	"github.com/geul-org/gogeul/pkg/resolve"
)

// Tree writes a static copy of the tree under root/.well-known/geul to
// dst/.well-known/geul (TreeFrom over the directory store).
func Tree(root, dst string) (Count, error) {
	return TreeFrom(resolve.Dir{Root: root}.Tree(), filepath.Join(root, ".well-known", "geul"), dst)
}
