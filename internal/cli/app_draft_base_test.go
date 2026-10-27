//ff:func feature=cli type=helper control=sequence
//ff:what TestAppDraftBase — an unpublished draft against an empty document; no DOC-ID, no data folder and a folder argument fail
package cli

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestAppDraftBase(t *testing.T) {
	h := newHarness(t)
	h.must("issuer", "new", "--name", "me")
	f := filepath.Join(h.data, "n.geul")
	h.must("new", "me/0x0000000000000005", "-o", f)
	if out := h.must("diff", f); !strings.HasPrefix(out, "+ SPEC-VERSION") {
		t.Fatalf("%q", out)
	}
	noID, _ := core.EncodeDocument([]core.Record{core.EncodeSpecVersion(2)}, nil)
	g := filepath.Join(h.data, "noid.geul")
	os.WriteFile(g, noID, 0o644)
	if _, err := h.run("", "diff", g); err == nil {
		t.Fatal("no DOC-ID: want error")
	}
	if _, err := h.run("", "diff", h.data); err == nil {
		t.Fatal("folder: want error")
	}
	h.cwd = t.TempDir()
	if _, err := h.run("", "diff", f); err == nil {
		t.Fatal("no data folder: want error")
	}
}
