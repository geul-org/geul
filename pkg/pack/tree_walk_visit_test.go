//ff:func feature=pack type=writer control=sequence
//ff:what TestTreeWalkVisit — files are kept as relative slash paths, directories skipped, walk errors returned
package pack

import (
	"errors"
	"io/fs"
	"path/filepath"
	"testing"
	"testing/fstest"
)

func TestTreeWalkVisit(t *testing.T) {
	m := fstest.MapFS{"a/b.geul": {}, "a": {Mode: fs.ModeDir}}
	tw := &treeWalk{root: "r"}
	f, _ := fs.Stat(m, "a/b.geul")
	d, _ := fs.Stat(m, "a")
	if err := tw.visit(filepath.Join("r", "a", "b.geul"), fs.FileInfoToDirEntry(f), nil); err != nil || tw.paths[0] != "a/b.geul" {
		t.Fatal(tw.paths, err)
	}
	if err := tw.visit("r/a", fs.FileInfoToDirEntry(d), nil); err != nil || len(tw.paths) != 1 {
		t.Fatal(tw.paths)
	}
	boom := errors.New("boom")
	if err := tw.visit("r/x", nil, boom); err != boom {
		t.Fatal(err)
	}
}
