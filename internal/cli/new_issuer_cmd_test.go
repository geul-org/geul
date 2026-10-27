//ff:func feature=cli type=command control=sequence
//ff:what TestNewIssuerCmd — new makes a 16-word Issuer, ls lists it, show prints its key
package cli

import (
	"strings"
	"testing"
)

func TestNewIssuerCmd(t *testing.T) {
	h := newHarness(t)
	out := h.must("issuer", "new", "--name", "me")
	name, id, _ := strings.Cut(strings.TrimSpace(out), " ")
	if name != "me" || len(id) != 66 {
		t.Fatal(out)
	}
	if got := h.must("issuer", "ls"); got != out {
		t.Fatal(got)
	}
	if got := h.must("issuer", "show", "me"); !strings.Contains(got, "issuer "+id) || !strings.Contains(got, "public-key 0x") {
		t.Fatal(got)
	}
	if got := h.must("issuer", "show", id); !strings.HasPrefix(got, "name me\n") {
		t.Fatal(got)
	}
}
