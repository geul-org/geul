//ff:func feature=cli type=helper control=sequence
//ff:what TestCheckBytes — one finding per schema Fact, under the given path; a clean document has none
package cli

import (
	"strings"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestCheckBytes(t *testing.T) {
	if fs := checkBytes("p", []byte{1, 2, 3}); len(fs) == 0 || fs[0].Path != "p" || !strings.Contains(fs[0].Message, "expected") {
		t.Fatal(fs)
	}
	doc, _ := core.EncodeDocument([]core.Record{core.EncodeSpecVersion(core.SpecVersion)}, nil)
	if fs := checkBytes("p", doc); len(fs) != 0 {
		t.Fatal(fs)
	}
}
