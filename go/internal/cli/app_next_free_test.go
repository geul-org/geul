//ff:func feature=cli type=helper control=sequence
//ff:what TestAppNextFree — above the published NEXT-LOCAL-ID after removals; a draft without DOC-ID counts only its own IDs
package cli

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestAppNextFree(t *testing.T) {
	h := newHarness(t)
	_, draft := h.doc()
	h.must("rm", draft, "#3")
	if out := h.must("add", draft, "ALIAS", "Q42", `"x"@en`); out != "#4\n" {
		t.Fatal(out)
	}
	noID, _ := core.EncodeDocument([]core.Record{core.EncodeSpecVersion(2)}, nil)
	f := filepath.Join(h.data, "noid.geul")
	os.WriteFile(f, noID, 0o644)
	if out := h.must("add", f, "LABEL", "Q1", `"a"`); out != "#1\n" {
		t.Fatal(out)
	}
}
