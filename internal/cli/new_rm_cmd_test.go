//ff:func feature=cli type=command control=sequence
//ff:what TestNewRmCmd — an Edge goes; one referred to stays unless --force
package cli

import (
	"strings"
	"testing"
)

func TestNewRmCmd(t *testing.T) {
	h := newHarness(t)
	_, draft := h.doc()
	h.must("rm", draft, "#3")
	if out := h.must("show", draft); strings.Contains(out, "ALIAS") {
		t.Fatal(out)
	}
	h.must("add", draft, "QUALIFIER", "P580", "#2", "GREGORIAN-DATE(1971-01-01)")
	if _, err := h.run("", "rm", draft, "#2"); err == nil {
		t.Fatal("referred to: want error")
	}
	h.must("rm", draft, "#2", "--force")
}
