//ff:func feature=pack type=writer control=sequence
//ff:what TestWriterWriteRecord — appends with an inventory entry; changed bytes, Get, codec and object-start errors fail
package pack

import (
	"errors"
	"testing"
)

func TestWriterWriteRecord(t *testing.T) {
	out := t.TempDir()
	never().MkdirAll(out + "/o")
	data := map[string][]byte{"a": []byte("abc")}
	w := testWriter(t, out, Batch{Put: []string{"a"}, Get: memGet(data)})
	w.plan()
	p := w.puts[0]
	if err := w.writeRecord(p); err != nil || p.loc.Off != 64 || len(w.inv.entries) != 1 {
		t.Fatal(err, p.loc)
	}
	data["a"] = []byte("abd")
	if err := w.writeRecord(p); err == nil {
		t.Fatal("changed")
	}
	data["a"] = []byte("abc")
	w.o.Codec = fakeCodec{err: errors.New("x")}
	if err := w.writeRecord(p); err == nil {
		t.Fatal("codec")
	}
	w.o.Codec = Stored{}
	w.abort()
	w.fs = &fsFuncs{mkdir: -1, create: 0, rename: -1}
	if err := w.writeRecord(p); err == nil {
		t.Fatal("start")
	}
}
