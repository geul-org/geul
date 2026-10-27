//ff:func feature=cli type=command control=sequence
//ff:what TestRmEdges — Edges go (HEAD Edges too), with --force even when referred to; a missing draft, a bad ID, a missing Edge and a reference without --force fail
package cli

import (
	"path/filepath"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
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
	h.must("add", draft, "SOURCE", "#0", "Q2013")
	h.must("add", draft, "QUALIFIER", "P813", "#5", "Q1")
	if err := rmEdges(draft, []string{"#5"}, false); err == nil {
		t.Fatal("HEAD SOURCE referred to by its QUALIFIER: want error")
	}
	if err := rmEdges(draft, []string{"#6", "#5"}, false); err != nil {
		t.Fatal(err)
	}
	if d, _ := readDraft(draft); headSourcesOf(d.head)[5] || d.head[len(d.head)-1].Kind != core.KindMeta {
		t.Fatalf("HEAD Edges removed: %+v", d.head)
	}
}
