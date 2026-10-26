//ff:func feature=cli type=helper control=sequence
//ff:what TestAppDiffPair — @v2 of the same document, two documents, a draft, one version against the newest, the last two
package cli

import (
	"strings"
	"testing"
)

func TestAppDiffPair(t *testing.T) {
	h := newHarness(t)
	addr, draft := h.doc()
	h.must("add", draft, "ALIAS", "Q42", `"D"@en`)
	if out := h.must("diff", draft); out != "+ #4 ALIAS Q42 \"D\"@en\n" {
		t.Fatalf("draft: %q", out)
	}
	if out := h.must("diff", addr+"@1", "@2"); !strings.Contains(out, "+ #3") {
		t.Fatalf("@v2: %q", out)
	}
	if out := h.must("diff", addr+"@1"); !strings.Contains(out, "+ #3") {
		t.Fatalf("against newest: %q", out)
	}
	if out := h.must("diff", addr); !strings.Contains(out, "+ #3") {
		t.Fatalf("last two: %q", out)
	}
	if out := h.must("diff", addr+"@1", addr+"@2"); !strings.Contains(out, "+ #3") {
		t.Fatalf("two addresses: %q", out)
	}
}
