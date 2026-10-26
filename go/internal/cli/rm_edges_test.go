//ff:func feature=cli type=command control=sequence
//ff:what TestRmEdges — Edges go, with --force even when referred to; a missing draft, a bad ID, a missing Edge and a reference without --force fail
package cli

import (
	"path/filepath"
	"testing"
)

func TestRmEdges(t *testing.T) {
	dir := t.TempDir()
	if err := rmEdges(filepath.Join(dir, "none"), []string{"#1"}, false); err == nil {
		t.Fatal("missing draft: want error")
	}
	p := brokenDraft(t, dir)
	if err := rmEdges(p, []string{"1"}, false); err == nil {
		t.Fatal("bad ID: want error")
	}
	if err := rmEdges(p, []string{"#9"}, false); err == nil {
		t.Fatal("missing Edge: want error")
	}
	h := newHarness(t)
	_, draft := h.doc()
	h.must("add", draft, "QUALIFIER", "P580", "#2", "Q1")
	if err := rmEdges(draft, []string{"#2"}, false); err == nil {
		t.Fatal("referred to: want error")
	}
	if err := rmEdges(draft, []string{"#2", "#3"}, true); err != nil {
		t.Fatal(err)
	}
	if d, _ := readDraft(draft); len(d.body) != 2 {
		t.Fatalf("%d left", len(d.body))
	}
}
