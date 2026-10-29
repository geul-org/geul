//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what printResults — one line per written pack set: generation, records, shared h/ paths, deletions, objects and bytes
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
		fmt.Fprintf(w, "geul: generation %d %s: %d records, %d shared h/, %d deleted, %d objects, %d bytes\n",
			r.Generation, r.Manifest, r.Records, r.Shared, r.Deleted, len(r.Objects), r.Bytes)
	}
}
