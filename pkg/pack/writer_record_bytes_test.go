//ff:func feature=pack type=writer control=sequence
//ff:what TestWriterRecordBytes — the sizes of records with their own bytes; shared h/ puts are not counted
package pack

import "testing"

func TestWriterRecordBytes(t *testing.T) {
	a := &put{path: "a", n: 10}
	w := &writer{inum: map[string]uint32{}, puts: []*put{a, {path: "h", n: 99, via: a}, {path: "b", n: 5}}}
	if n := w.recordBytes(); n != 15 {
		t.Fatal(n)
	}
}
