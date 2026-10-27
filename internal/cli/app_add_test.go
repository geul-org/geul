//ff:func feature=cli type=command control=sequence
//ff:what TestAppAdd — no Edge, a missing draft, a line that does not read, a repeated label and an Edge breaking the spec fail and write nothing; labels (#n, #name, also forward) map to new IDs, other #n stay
package cli

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestAppAdd(t *testing.T) {
	h := newHarness(t)
	_, draft := h.doc()
	before, _ := os.ReadFile(draft)
	if _, err := h.run("", "add", draft); err == nil {
		t.Fatal("no Edge: want error")
	}
	if _, err := h.run("", "add", filepath.Join(h.data, "none"), "LABEL", "Q1", `"a"`); err == nil {
		t.Fatal("missing draft: want error")
	}
	if _, err := h.run("ALIAS Q1 \"a\"\nNOPE Q1\n", "add", draft, "-i", "-"); err == nil {
		t.Fatal("bad line: want error")
	}
	if _, err := h.run("#a ALIAS Q1 \"a\"\n#a ALIAS Q1 \"b\"\n", "add", draft, "-i", "-"); err == nil {
		t.Fatal("a label used twice: want error")
	}
	if _, err := h.run("", "add", draft, "FIELD", "Q1"); err == nil {
		t.Fatal("FIELD without its layout breaks the spec: want error")
	}
	if after, _ := os.ReadFile(draft); string(after) != string(before) {
		t.Fatal("draft changed")
	}
	out, err := h.run("#1 TRIPLE P69 Q42 Q1 normal\nQUALIFIER P580 #1 #2\n", "add", draft, "-i", "-")
	if err != nil || out != "#4\n#5\n" {
		t.Fatalf("%q %v", out, err)
	}
	if show := h.must("show", draft); !strings.Contains(show, "#5 QUALIFIER P580 #4 #2") {
		t.Fatal(show)
	}
	out, err = h.run("QUALIFIER P582 #edu #2\n#edu TRIPLE P69 Q42 Q2 normal\n", "add", draft, "-i", "-")
	if err != nil || out != "#6\n#7\n" {
		t.Fatalf("%q %v", out, err)
	}
	if show := h.must("show", draft); !strings.Contains(show, "#6 QUALIFIER P582 #7 #2") {
		t.Fatal(show)
	}
}
