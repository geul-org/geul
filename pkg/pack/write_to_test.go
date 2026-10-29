//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what TestWriteTo — no change makes no generation; a broken latest.txt, Get errors, mkdir failures, data and finish failures are returned
package pack

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestWriteTo(t *testing.T) {
	ctx := context.Background()
	_, out := built(t, small())
	paths, _ := listTree("testdata/archive")
	get := func(p string) ([]byte, error) { return fixtureFile(t, p), nil }
	res, err := writeTo(ctx, out, Batch{Put: paths, Full: true, Get: get}, small(), never())
	if err != nil || res.Changed || res.Generation != 1 {
		t.Fatal(res, err)
	}
	b := Batch{Put: []string{"n.geul"}, Get: memGet(map[string][]byte{"n.geul": {1}})}
	for i, fs := range []*fsFuncs{{mkdir: 0, create: -1, rename: -1}, {mkdir: 1, create: -1, rename: -1}, {mkdir: -1, create: 0, rename: -1}, {mkdir: -1, create: 1, rename: -1}} {
		if _, err := writeTo(ctx, out, b, small(), fs); err == nil {
			t.Fatal(i)
		}
	}
	if _, err := writeTo(ctx, out, Batch{Put: []string{"none"}, Get: memGet(nil)}, small(), never()); err == nil {
		t.Fatal("get")
	}
	bad := filepath.Join(t.TempDir(), "bad")
	os.MkdirAll(bad, 0o755)
	os.WriteFile(filepath.Join(bad, latestFile), []byte("x"), 0o644)
	if _, err := writeTo(ctx, bad, b, small(), never()); !errors.Is(err, ErrCorrupt) {
		t.Fatal(err)
	}
	if res, err := writeTo(ctx, out, b, small(), never()); err != nil || res.Generation != 2 {
		t.Fatal(res, err)
	}
}
