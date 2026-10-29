//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what TestArchiveUnpack — every path back as the original file; an existing tree, listing errors and unreadable paths fail
package pack

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestArchiveUnpack(t *testing.T) {
	a, _ := testArchive(t)
	dst := t.TempDir()
	if n, err := a.Unpack(dst); err != nil || n != 18 {
		t.Fatal(n, err)
	}
	paths, _ := listTree(dst)
	for _, p := range paths {
		b, _ := os.ReadFile(filepath.Join(dst, ".well-known", "geul", filepath.FromSlash(p)))
		if !bytes.Equal(b, fixtureFile(t, p)) {
			t.Fatal(p)
		}
	}
	if _, err := a.Unpack(dst); err == nil {
		t.Fatal("exists")
	}
	os.Remove(objPath(t, a.sets[0].r, a.sets[0].dir, kindData))
	if _, err := a.Unpack(t.TempDir()); err == nil {
		t.Fatal("unreadable")
	}
	b, _ := testArchive(t)
	b.sets[0].r.m.pLayers = []layer{{blk: block{kind: blkPInner, keys: []string{"a"}, locs: []Loc{{Obj: 99}}}}}
	if _, err := b.Unpack(t.TempDir()); err == nil {
		t.Fatal("listing")
	}
}
