//ff:func feature=cli type=command control=sequence
//ff:what TestNewBlameCmd — each Edge with the version adding it and the one changing it
package cli

import "testing"

func TestNewBlameCmd(t *testing.T) {
	h := newHarness(t)
	addr, _ := h.doc()
	want := "1\t1\t#1 LABEL Q42 \"Douglas Adams\"@en\n1\t2\t#2 TRIPLE P31 Q42 Q36180 normal\n2\t2\t#3 ALIAS Q42 \"DNA\"@en\n"
	if out := h.must("blame", addr); out != want {
		t.Fatalf("%q", out)
	}
}
