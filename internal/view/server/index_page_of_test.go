//ff:func feature=server type=helper control=sequence
//ff:what TestIndexPageOf — parent link below the top, none at the top; directories labelled at any level, files never; a missing directory fails
package server

import (
	"path/filepath"
	"testing"
)

func TestIndexPageOf(t *testing.T) {
	top := filepath.Join(treeFixture(t), ".well-known", "geul")
	if p, err := indexPageOf(filepath.Join(top, "h"), "/.well-known/geul/h/", map[string]string{"a.geul": "x"}); err != nil || p.Up != "/.well-known/geul/" {
		t.Fatalf("sub: %+v %v", p, err)
	}
	if p, err := indexPageOf(top, "/.well-known/geul/", map[string]string{"h": "hub"}); err != nil || p.Up != "" || len(p.Entries) != 1 || p.Entries[0] != (indexEntry{Name: "hub/", URL: "/.well-known/geul/h/", ID: "h"}) {
		t.Fatalf("top: %+v %v", p, err)
	}
	if p, _ := indexPageOf(top, "/.well-known/geul/0x0000/", map[string]string{"h": "doc"}); p.Entries[0].Name != "doc/" {
		t.Fatalf("below top: %+v", p)
	}
	if _, err := indexPageOf(filepath.Join(top, "none"), "/x/", nil); err == nil {
		t.Fatal("missing: want error")
	}
}
