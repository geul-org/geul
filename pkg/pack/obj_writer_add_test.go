//ff:func feature=pack type=helper control=sequence
//ff:what TestObjWriterAdd — returns the item's offset and counts it
package pack

import "testing"

func TestObjWriterAdd(t *testing.T) {
	f, _ := osFS{}.CreateTemp(t.TempDir())
	w := &objWriter{f: f, off: 64}
	if w.add([]byte("abc")) != 64 || w.add([]byte("d")) != 67 || w.count != 2 {
		t.Fatal(w)
	}
}
