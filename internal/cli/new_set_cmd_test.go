//ff:func feature=cli type=command control=sequence
//ff:what TestNewSetCmd — an Edge's fields change in place under its ID; too few arguments fail
package cli

import (
	"strings"
	"testing"
)

func TestNewSetCmd(t *testing.T) {
	h := newHarness(t)
	_, draft := h.doc()
	h.must("set", draft, "#1", "Q42", `"D. Adams"@en`)
	if out := h.must("show", draft); !strings.Contains(out, "BODY\n#1 LABEL Q42 \"D. Adams\"@en\n") {
		t.Fatal(out)
	}
	if _, err := h.run("", "set", draft); err == nil {
		t.Fatal("want error")
	}
}
