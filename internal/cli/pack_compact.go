//ff:func feature=cli type=command control=sequence
//ff:what packCompact — rewrite every set of a pack archive into fresh objects (one P and one H layer each)
package cli

import (
	"context"
	"io"
)

// packCompact compacts dir's pack archive; older objects stay on disk until removed.
func packCompact(dir string, w io.Writer) error {
	a, err := openArchive(dir)
	if err != nil {
		return err
	}
	rs, err := a.Compact(context.Background())
	if err != nil {
		return err
	}
	printResults(w, rs)
	return nil
}
