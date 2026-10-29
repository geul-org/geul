//ff:func feature=pack type=writer control=sequence
//ff:what TestWriterReadRecord — the encoded record; changed bytes, Get and codec errors fail
package pack

import (
	"errors"
	"testing"
)

func TestWriterReadRecord(t *testing.T) {
	data := map[string][]byte{"a": []byte("abc")}
	w := testWriter(t, t.TempDir(), Batch{Put: []string{"a"}, Get: memGet(data)})
	w.plan()
	p := w.puts[0]
	if e, err := w.readRecord(p); err != nil || e.enc != encRaw || string(e.b) != "abc" {
		t.Fatal(e, err)
	}
	data["a"] = []byte("abd")
	if _, err := w.readRecord(p); err == nil {
		t.Fatal("changed")
	}
	delete(data, "a")
	if _, err := w.readRecord(p); err == nil {
		t.Fatal("get")
	}
	data["a"] = []byte("abc")
	w.o.Codec = fakeCodec{err: errors.New("x")}
	if _, err := w.readRecord(p); err == nil {
		t.Fatal("codec")
	}
}
