//ff:func feature=resolve type=resolver control=sequence
//ff:what TestDirWrite — creates directories; fails when a parent is a regular file
package resolve

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDirWrite(t *testing.T) {
	root := t.TempDir()
	d := Dir{Root: root}
	if err := d.write("x/y/z.geul", []byte("z")); err != nil {
		t.Fatal(err)
	}
	if b, _ := os.ReadFile(filepath.Join(root, ".well-known", "geul", "x", "y", "z.geul")); string(b) != "z" {
		t.Fatalf("got %q", b)
	}
	if err := d.write("x/y/z.geul/w", []byte("w")); err == nil {
		t.Fatal("file as parent: want error")
	}
}
