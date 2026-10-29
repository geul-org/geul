//ff:func feature=cli type=command control=sequence
//ff:what packBuild — pack a directory archive into pack sets under another data folder (one set per Issuer, or one set)
package cli

import (
	"context"
	"fmt"
	"io"
	"path/filepath"

	"github.com/geul-org/gogeul/pkg/pack"
)

// packBuild converts the directory archive src holds (or is) into pack sets
// under out/.well-known/geul/; an existing pack archive there gets its next
// generations. out must be another folder. o.Codec stores the records.
func packBuild(src, out string, single bool, o pack.BuildOptions, w io.Writer) error {
	root, err := treeRoot(src)
	if err != nil {
		return err
	}
	wk := filepath.Join(out, ".well-known", "geul")
	ra, _ := filepath.Abs(root)
	oa, _ := filepath.Abs(out)
	if ra == oa {
		return fmt.Errorf("%s: pack into another folder, then replace the tree", out)
	}
	if pack.IsPacked(filepath.Join(root, ".well-known", "geul")) {
		return fmt.Errorf("%s: already a pack archive", src)
	}
	var rs []pack.BuildResult
	if single {
		var r pack.BuildResult
		r, err = pack.Build(context.Background(), root, wk, o)
		rs = append(rs, r)
	} else {
		rs, err = pack.Split(context.Background(), root, wk, o)
	}
	if err != nil {
		return err
	}
	printResults(w, rs)
	return nil
}
