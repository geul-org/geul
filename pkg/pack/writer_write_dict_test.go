//ff:func feature=pack type=writer control=sequence
//ff:what TestWriterWriteDict — the dictionary opens the first data pack: inventory kind 4, a DICTS entry with its SHA-256; a failed object start fails
package pack

import (
	"crypto/sha256"
	"testing"
)

func TestWriterWriteDict(t *testing.T) {
	out := t.TempDir()
	never().MkdirAll(out + "/o")
	w := testWriter(t, out, Batch{})
	d := &Dict{ID: FirstDictID, Bytes: []byte("dictionary bytes")}
	if err := w.writeDict(d); err != nil || w.dict != d || w.res.Dict != FirstDictID {
		t.Fatal(err)
	}
	e := w.inv.entries[0]
	if e.kind != invDict || e.ref != FirstDictID || e.off != headerSize || e.n != 16 {
		t.Fatal(e)
	}
	if de := w.dicts[0]; de.id != FirstDictID || de.loc.Off != headerSize || de.loc.Len != 16 || de.hash != sha256.Sum256(d.Bytes) {
		t.Fatal(de)
	}
	w.abort()
	w.fs = &fsFuncs{mkdir: -1, create: 0, rename: -1}
	if err := w.writeDict(d); err == nil {
		t.Fatal("start")
	}
}
