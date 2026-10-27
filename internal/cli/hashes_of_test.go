//ff:func feature=cli type=helper control=sequence
//ff:what TestHashesOf — a published document's hash is in the set; an unpublished document and no DOC-ID give none
package cli

import (
	"os"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/resolve"
)

func TestHashesOf(t *testing.T) {
	h := newHarness(t)
	_, draft := h.doc()
	b, _ := os.ReadFile(draft)
	d := resolve.Dir{Root: h.data}
	issuer, doc, _ := docIDOf(b)
	lg, _ := d.Log(issuer, doc)
	if got := hashesOf(d, b); len(got) != 2 || !got[lg.Entries[1].Hash] {
		t.Fatal(got)
	}
	other, _ := core.EncodeDocument([]core.Record{core.EncodeSpecVersion(2)}, nil)
	if got := hashesOf(d, other); len(got) != 0 {
		t.Fatal(got)
	}
}
