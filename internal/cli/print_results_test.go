//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what TestPrintResults — nothing written, unchanged and written generations, zstd records and a new dictionary
package cli

import (
	"bytes"
	"strings"
	"testing"

	"github.com/geul-org/gogeul/pkg/pack"
)

func TestPrintResults(t *testing.T) {
	var b bytes.Buffer
	printResults(&b, nil)
	printResults(&b, []pack.BuildResult{{Generation: 1}, {Generation: 2, Changed: true, Manifest: "m", Records: 3, Objects: []string{"o"}}, {Generation: 3, Changed: true, Records: 4, Compressed: 4, Dict: 32768}})
	for _, want := range []string{"nothing to write", "generation 1 unchanged", "generation 2 m: 3 records (0 zstd), 0 shared h/, 0 deleted, 1 objects", "4 records (4 zstd)", "bytes, dictionary 32768\n"} {
		if !strings.Contains(b.String(), want) {
			t.Fatal(b.String())
		}
	}
}
