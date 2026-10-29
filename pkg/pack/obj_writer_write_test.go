//ff:func feature=pack type=helper control=sequence
//ff:what TestObjWriterWrite — appends and advances; after a failure it only advances
package pack

import (
	"testing"
)

func TestObjWriterWrite(t *testing.T) {
	f, _ := osFS{}.CreateTemp(t.TempDir())
	bf := &badFile{objFile: f}
	w := &objWriter{f: bf}
	w.write([]byte("ab"))
	if w.off != 2 || w.err != nil {
		t.Fatal(w)
	}
	bf.write = true
	w.write([]byte("c"))
	w.write([]byte("d"))
	if w.off != 4 || w.err == nil {
		t.Fatal(w)
	}
}
