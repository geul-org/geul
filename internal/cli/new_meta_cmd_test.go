//ff:func feature=cli type=command control=sequence
//ff:what TestNewMetaCmd — a Meta is set, replaced in place and removed
package cli

import (
	"strings"
	"testing"
)

func TestNewMetaCmd(t *testing.T) {
	h := newHarness(t)
	_, draft := h.doc()
	h.must("meta", draft, "MESSAGE", `"one"`)
	h.must("meta", draft, "message", `"two"`)
	out := h.must("show", draft)
	if strings.Count(out, "MESSAGE") != 1 || !strings.Contains(out, `MESSAGE "two"`) {
		t.Fatal(out)
	}
	h.must("meta", draft, "MESSAGE", "--unset")
	if out := h.must("show", draft); strings.Contains(out, "MESSAGE") {
		t.Fatal(out)
	}
}
