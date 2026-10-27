//ff:func feature=server type=helper control=sequence
//ff:what TestIndexEntries — directories and .geul files in order, other files skipped; a missing directory fails
package server

import (
	"path/filepath"
	"testing"
)

func TestIndexEntries(t *testing.T) {
	root := treeFixture(t)
	es, err := indexEntries(filepath.Join(root, ".well-known", "geul", "h"), "/b/")
	if err != nil || len(es) != 2 || es[0] != (indexEntry{Name: "a.geul", URL: "/b/a"}) || es[1].Name != "bad.geul" {
		t.Fatalf("files: %+v %v", es, err)
	}
	es, _ = indexEntries(filepath.Join(root, ".well-known", "geul"), "/c/")
	if len(es) != 1 || es[0] != (indexEntry{Name: "h/", URL: "/c/h/"}) {
		t.Fatalf("dirs: %+v", es)
	}
	if _, err := indexEntries(filepath.Join(root, "none"), "/"); err == nil {
		t.Fatal("missing: want error")
	}
}
