//ff:func feature=resolve type=resolver control=sequence
//ff:what TestDirFile — documents map under .well-known/geul, list paths to the versions file
package resolve

import (
	"path/filepath"
	"testing"
)

func TestDirFile(t *testing.T) {
	d := Dir{Root: "/r"}
	if f := d.file("0x0000/0x0000000000000001/0x00000001.geul"); f != filepath.FromSlash("/r/.well-known/geul/0x0000/0x0000000000000001/0x00000001.geul") {
		t.Fatalf("doc: %q", f)
	}
	if f := d.file("0x0000/0x0000000000000001/"); f != filepath.FromSlash("/r/.well-known/geul/0x0000/0x0000000000000001/versions") {
		t.Fatalf("list: %q", f)
	}
}
