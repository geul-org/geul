//ff:func feature=pack type=writer control=sequence
//ff:what TestWriterPlaceRecord — appends with an inventory entry, counts zstd records, opens a new pack when full; a failed object start fails
package pack

import "testing"

func TestWriterPlaceRecord(t *testing.T) {
	out := t.TempDir()
	never().MkdirAll(out + "/o")
	w := testWriter(t, out, Batch{})
	w.o.PackSize = 64 + 10
	p := &put{path: "a", ref: parsePath("a")}
	if err := w.placeRecord(p, encoded{enc: encZstd, b: []byte("12345")}); err != nil || p.loc.Off != 64 || p.loc.Enc != encZstd || w.res.Compressed != 1 {
		t.Fatal(err, p.loc)
	}
	q := &put{path: "b", ref: parsePath("b")}
	if err := w.placeRecord(q, encoded{b: []byte("123456789")}); err != nil || q.loc.Obj != 1 || w.res.Records != 2 || w.res.Compressed != 1 {
		t.Fatal(err, q.loc)
	}
	w.abort()
	w.fs = &fsFuncs{mkdir: -1, create: 0, rename: -1}
	if err := w.placeRecord(q, encoded{b: []byte("x")}); err == nil {
		t.Fatal("start")
	}
}
