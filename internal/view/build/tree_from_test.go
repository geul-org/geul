//ff:func feature=build type=builder control=sequence
//ff:what TestTreeFrom — a pack archive bakes to the same files as its directory tree; output inside the source and an unreadable tree fail
package build

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/geul-org/gogeul/pkg/pack"
	"github.com/geul-org/gogeul/pkg/resolve"
)

func TestTreeFrom(t *testing.T) {
	root := fixture(t)
	wk := filepath.Join(t.TempDir(), ".well-known", "geul")
	if _, err := pack.Build(context.Background(), root, wk, pack.BuildOptions{}); err != nil {
		t.Fatal(err)
	}
	a, _ := pack.OpenDir(context.Background(), wk, pack.BuildOptions{})
	dirOut, packOut := t.TempDir(), t.TempDir()
	c1, err1 := TreeFrom(resolve.Dir{Root: root}.Tree(), filepath.Join(root, ".well-known", "geul"), dirOut)
	c2, err2 := TreeFrom(a, wk, packOut)
	if err1 != nil || err2 != nil || c1 != c2 {
		t.Fatal(c1, c2, err1, err2)
	}
	want, _ := os.ReadFile(filepath.Join(dirOut, ".well-known", "geul", "h", "a"))
	got, _ := os.ReadFile(filepath.Join(packOut, ".well-known", "geul", "h", "a"))
	if string(got) != string(want) || len(got) == 0 {
		t.Fatal("pages differ")
	}
	if _, err := TreeFrom(a, wk, filepath.Dir(filepath.Dir(wk))); err == nil {
		t.Fatal("inside the source")
	}
	if _, err := TreeFrom(resolve.Dir{Root: t.TempDir()}.Tree(), "src", t.TempDir()); err == nil {
		t.Fatal("unreadable tree")
	}
}
