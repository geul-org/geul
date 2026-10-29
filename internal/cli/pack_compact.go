//ff:func feature=cli type=command control=sequence
//ff:what packCompact — rewrite every set of a pack archive into fresh objects (one P and one H layer each)
package cli

import (
	"context"
	"io"

	"github.com/geul-org/gogeul/pkg/pack"
)

// packCompact compacts dir's pack archive with o (its Codec nil keeps each
// set's codec); older objects stay on disk until removed.
func packCompact(dir string, o pack.BuildOptions, w io.Writer) error {
	a, err := openArchive(dir, o)
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
