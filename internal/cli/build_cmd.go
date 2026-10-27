//ff:func feature=cli type=command control=sequence
//ff:what buildCmd — find the tree, bake it under the output root and report what was written
package cli

import (
	"fmt"
	"io"
	"path/filepath"

	"github.com/geul-org/gogeul/internal/view/build"
)

// buildCmd bakes the tree arg holds (or is) under dst and prints a summary.
func buildCmd(arg, dst string, out io.Writer) error {
	root, err := treeRoot(arg)
	if err != nil {
		return err
	}
	c, err := build.Tree(root, dst)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(out, "geul: %d documents, %d directories → %s\n", c.Docs, c.Dirs, filepath.Join(dst, ".well-known", "geul"))
	return err
}
