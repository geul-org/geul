//ff:func feature=pack type=writer control=sequence
//ff:what TestWriterWriteBatch — records placed in order; a read error and a place error stop it
package pack

import "testing"

func TestWriterWriteBatch(t *testing.T) {
	out := t.TempDir()
	never().MkdirAll(out + "/o")
	paths, get := synth(20)
	w := testWriter(t, out, Batch{Put: paths, Get: get})
	w.o.Codec = &Zstd{Workers: 4, DictSize: -1}
	w.plan()
	rs := w.records()
	if err := w.writeBatch(rs); err != nil || w.res.Records != 20 || w.res.Compressed != 20 || rs[19].loc.Off <= rs[0].loc.Off {
		t.Fatal(err, w.res)
	}
	w.b.Get = memGet(nil)
	if err := w.writeBatch(rs); err == nil {
		t.Fatal("read")
	}
	w.b.Get = get
	w.abort()
	w.fs = &fsFuncs{mkdir: -1, create: 0, rename: -1}
	if err := w.writeBatch(rs[:1]); err == nil {
		t.Fatal("place")
	}
}
