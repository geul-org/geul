//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what TestUnpack — every path back as a file equal to the original; an existing tree, index errors and unreadable records fail
package pack

import (
	"bytes"
	"context"
	"os"
	"path/filepath"
	"testing"
)

func TestUnpack(t *testing.T) {
	ctx := context.Background()
	_, out := built(t, small())
	r := opened(t, out)
	dst := t.TempDir()
	n, err := Unpack(ctx, r, dst)
	if err != nil || n != 18 {
		t.Fatal(n, err)
	}
	paths, _ := listTree(dst)
	for _, p := range paths {
		b, _ := os.ReadFile(filepath.Join(dst, ".well-known", "geul", filepath.FromSlash(p)))
		if !bytes.Equal(b, fixtureFile(t, p)) {
			t.Fatal(p)
		}
	}
	if _, err := Unpack(ctx, r, dst); err == nil {
		t.Fatal("exists")
	}
	os.Remove(objPath(t, r, out, kindData))
	if _, err := Unpack(ctx, r, t.TempDir()); err == nil {
		t.Fatal("records")
	}
	r.m.pLayers = []layer{{blk: block{kind: blkPInner, keys: []string{"a"}, locs: []Loc{{Obj: 99}}}}}
	if _, err := Unpack(ctx, r, t.TempDir()); err == nil {
		t.Fatal("index")
	}
}
