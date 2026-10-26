//ff:func feature=cli type=helper control=sequence
//ff:what TestAppReadDoc — a file as is, an address from the data folder; a bad address and an address without a data folder fail
package cli

import (
	"testing"
)

func TestAppReadDoc(t *testing.T) {
	h := newHarness(t)
	addr, draft := h.doc()
	if a, b := h.must("hash", draft), h.must("hash", addr); a != b {
		t.Fatalf("%s %s", a, b)
	}
	if _, err := h.run("", "show", "nobody/x"); err == nil {
		t.Fatal("bad address: want error")
	}
	h.cwd = t.TempDir()
	if _, err := h.run("", "show", "0x0000/0x0000000000000001"); err == nil {
		t.Fatal("no data folder: want error")
	}
}
