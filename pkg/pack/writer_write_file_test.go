//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what TestWriterWriteFile — atomic publish; create, write, close and rename failures leave no file
package pack

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriterWriteFile(t *testing.T) {
	out := t.TempDir()
	w := &writer{out: out, fs: never()}
	if err := w.writeFile(latestFile, []byte("x")); err != nil {
		t.Fatal(err)
	}
	if b, _ := os.ReadFile(filepath.Join(out, latestFile)); string(b) != "x" {
		t.Fatal(string(b))
	}
	cases := []*fsFuncs{{create: 0}, {create: -1, rename: 0}}
	for _, flag := range []badFile{{write: true}, {close: true}} {
		flag := flag
		cases = append(cases, &fsFuncs{create: -1, rename: -1, wrap: func(f objFile) objFile { flag.objFile = f; return &flag }})
	}
	for i, fs := range cases {
		w.fs = fs
		if err := w.writeFile("y", []byte("y")); err == nil {
			t.Fatal(i)
		}
	}
	if _, err := os.Stat(filepath.Join(out, "y")); err == nil {
		t.Fatal("y written")
	}
}
