//ff:func feature=pack type=helper control=sequence
//ff:what TestDirFileReadDir — in pages of n, io.EOF after the last; n ≤ 0 returns the rest
package pack

import (
	"io"
	"testing"
)

func TestDirFileReadDir(t *testing.T) {
	d := &dirFile{rest: []dirEntry{{name: "a"}, {name: "b"}, {name: "c"}}}
	if es, err := d.ReadDir(2); err != nil || len(es) != 2 || es[1].Name() != "b" {
		t.Fatal(es, err)
	}
	if es, err := d.ReadDir(0); err != nil || len(es) != 1 {
		t.Fatal(es, err)
	}
	if _, err := d.ReadDir(1); err != io.EOF {
		t.Fatal(err)
	}
}
