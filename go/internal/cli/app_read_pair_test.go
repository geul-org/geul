//ff:func feature=cli type=helper control=sequence
//ff:what TestAppReadPair — both read (the same version twice has no changes); a missing base or next fails
package cli

import "testing"

func TestAppReadPair(t *testing.T) {
	h := newHarness(t)
	addr, _ := h.doc()
	if _, err := h.run("", "diff", addr+"@9", addr+"@1"); err == nil {
		t.Fatal("missing base: want error")
	}
	if _, err := h.run("", "diff", addr+"@1", addr+"@9"); err == nil {
		t.Fatal("missing next: want error")
	}
	if out := h.must("diff", addr+"@1", addr+"@1"); out != "" {
		t.Fatal(out)
	}
}
