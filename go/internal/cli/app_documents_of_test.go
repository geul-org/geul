//ff:func feature=cli type=command control=sequence
//ff:what TestAppDocumentsOf — documents of a local key's Issuer; an unknown name fails
package cli

import (
	"strings"
	"testing"
)

func TestAppDocumentsOf(t *testing.T) {
	h := newHarness(t)
	h.doc()
	if out := h.must("ls", "me"); !strings.HasPrefix(out, "0x0000000000000001") {
		t.Fatal(out)
	}
	if _, err := h.run("", "ls", "nobody"); err == nil {
		t.Fatal("want error")
	}
}
