//ff:func feature=cli type=command control=sequence
//ff:what TestNewGrepCmd — matching lines of every document with their address; no pattern fails
package cli

import (
	"strings"
	"testing"
)

func TestNewGrepCmd(t *testing.T) {
	h := newHarness(t)
	h.doc()
	out := h.must("grep", "DNA")
	if !strings.HasSuffix(out, "/0x0000000000000001: #3 ALIAS Q42 \"DNA\"@en\n") || strings.Count(out, "\n") != 1 {
		t.Fatalf("%q", out)
	}
	if _, err := h.run("", "grep"); err == nil {
		t.Fatal("want error")
	}
}
