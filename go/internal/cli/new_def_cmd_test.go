//ff:func feature=cli type=command control=sequence
//ff:what TestNewDefCmd — a REFERENCE-NODE with two new IDs, the node printed; one argument fails
package cli

import (
	"strings"
	"testing"
)

func TestNewDefCmd(t *testing.T) {
	h := newHarness(t)
	_, draft := h.doc()
	if out := h.must("def", draft, "<http://wikiba.se/ontology#propertyType>"); out != "#5\n" {
		t.Fatal(out)
	}
	if out := h.must("show", draft); !strings.Contains(out, "#4 REFERENCE-NODE #5 <http://wikiba.se/ontology#propertyType>") {
		t.Fatal(out)
	}
	if _, err := h.run("", "def", draft); err == nil {
		t.Fatal("no object: want error")
	}
}
