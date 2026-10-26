//ff:func feature=cli type=command control=sequence
//ff:what TestNewAddCmd — one Edge from arguments, a batch from a file with labels; a missing list file fails
package cli

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestNewAddCmd(t *testing.T) {
	h := newHarness(t)
	_, draft := h.doc()
	if out := h.must("add", draft, "ALIAS", "Q42", `"D"@en`); out != "#4\n" {
		t.Fatal(out)
	}
	list := filepath.Join(h.data, "list.txt")
	os.WriteFile(list, []byte("// comment\n#1 TRIPLE P69 Q42 Q1 normal\n\nQUALIFIER P580 #1 GREGORIAN-DATE(1971-01-01)\n"), 0o644)
	if out := h.must("add", draft, "-i", list); out != "#5\n#6\n" {
		t.Fatal(out)
	}
	if out := h.must("show", draft); !strings.Contains(out, "#6 QUALIFIER P580 #5 GREGORIAN-DATE(1971-01-01)") {
		t.Fatal(out)
	}
	if _, err := h.run("", "add", draft, "-i", filepath.Join(h.data, "none")); err == nil {
		t.Fatal("missing list: want error")
	}
}
