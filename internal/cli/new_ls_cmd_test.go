//ff:func feature=cli type=command control=sequence
//ff:what TestNewLsCmd — ls lists Issuers, an Issuer's documents with names, as JSON; extra arguments fail
package cli

import (
	"strings"
	"testing"
)

func TestNewLsCmd(t *testing.T) {
	h := newHarness(t)
	addr, _ := h.doc()
	h.must("name", addr, "adams")
	if out := h.must("ls"); !strings.HasPrefix(out, "0x") {
		t.Fatal(out)
	}
	if out := h.must("ls", "me"); out != "0x0000000000000001 adams\n" {
		t.Fatal(out)
	}
	if out := h.must("ls", "me", "--format", "json"); !strings.Contains(out, `"adams"`) {
		t.Fatal(out)
	}
	if _, err := h.run("", "ls", "a", "b"); err == nil {
		t.Fatal("two args: want error")
	}
}
