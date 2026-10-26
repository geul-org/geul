//ff:func feature=cli type=command control=sequence
//ff:what TestNewDiffCmd — diff of two versions shows the changed Edge and the added one; no arguments fail
package cli

import (
	"testing"
)

func TestNewDiffCmd(t *testing.T) {
	h := newHarness(t)
	addr, _ := h.doc()
	want := "~ #2 TRIPLE P31 Q42 Q5 normal\n→ #2 TRIPLE P31 Q42 Q36180 normal\n+ #3 ALIAS Q42 \"DNA\"@en\n"
	if out := h.must("diff", addr+"@1", "@2"); out != want {
		t.Fatalf("%q", out)
	}
	if _, err := h.run("", "diff"); err == nil {
		t.Fatal("want error")
	}
}
