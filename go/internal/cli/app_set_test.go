//ff:func feature=cli type=command control=iteration dimension=1
//ff:what TestAppSet — a bad ID, a missing draft, a missing Edge, fields that do not read and a draft breaking the spec fail
package cli

import (
	"path/filepath"
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
}
