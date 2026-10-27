//ff:func feature=cli type=command control=iteration dimension=1
//ff:what TestAppSet — a bad ID, a missing draft, a missing Edge, fields that do not read and a draft breaking the spec fail; a HEAD SOURCE is set in place, not onto a record
package cli

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestAppSet(t *testing.T) {
	h := newHarness(t)
	_, draft := h.doc()
	cases := [][]string{{draft, "1", "Q1"}, {filepath.Join(h.data, "none"), "#1", "Q1"}, {draft, "#9", "Q1"}, {draft, "#1", "word"}, {brokenDraft(t, h.data), "#1", "Q2"}}
	for _, c := range cases {
		if _, err := h.run("", append([]string{"set"}, c...)...); err == nil {
			t.Errorf("%v: want error", c)
		}
	}
	h.must("add", draft, "SOURCE", "#0", "Q2013")
	h.must("set", draft, "#4", "#0", "Q1")
	if head, _, _ := strings.Cut(h.must("show", draft), "BODY"); !strings.Contains(head, "#4 SOURCE #0 Q1") {
		t.Fatalf("HEAD SOURCE set in place: %s", head)
	}
	if _, err := h.run("", "set", draft, "#4", "#2", "Q1"); err == nil {
		t.Fatal("a HEAD SOURCE turned onto a record breaks the HEAD: want error")
	}
}
