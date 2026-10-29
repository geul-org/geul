//ff:func feature=pack type=writer control=sequence
//ff:what TestWriterStartData — closes the open pack first; a close failure is returned
package pack

import "testing"

func TestWriterStartData(t *testing.T) {
	out := t.TempDir()
	never().MkdirAll(out + "/o")
	w := &writer{out: out, fs: never(), o: small().withDefaults()}
	w.o.Codec = Stored{}
	if err := w.startData(); err != nil || w.cur == nil {
		t.Fatal(err)
	}
	if err := w.startData(); err != nil || len(w.objs) != 1 || w.curNum != 1 {
		t.Fatal(err)
	}
	w.fs = &fsFuncs{create: -1, rename: 0}
	if err := w.startData(); err == nil {
		t.Fatal("close")
	}
}
