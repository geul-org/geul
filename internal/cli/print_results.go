//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what printResults — one line per written pack set: generation, records (zstd ones), shared h/ paths, deletions, objects, bytes, new dictionary
package cli

import (
	"fmt"
	"io"

	"github.com/geul-org/gogeul/pkg/pack"
)

// printResults reports pack writes, or that nothing changed.
func printResults(w io.Writer, rs []pack.BuildResult) {
	if len(rs) == 0 {
		fmt.Fprintln(w, "geul: nothing to write")
	}
	for _, r := range rs {
		if !r.Changed {
			fmt.Fprintf(w, "geul: generation %d unchanged\n", r.Generation)
			continue
		}
		fmt.Fprintf(w, "geul: generation %d %s: %d records (%d zstd), %d shared h/, %d deleted, %d objects, %d bytes%s\n",
			r.Generation, r.Manifest, r.Records, r.Compressed, r.Shared, r.Deleted, len(r.Objects), r.Bytes, dictNote(r.Dict))
	}
}
