//ff:func feature=pack type=writer control=sequence
//ff:what TestWriterCloseData — deletions alone open a pack for their inventory; a start failure is returned
package pack

import "testing"

func TestWriterCloseData(t *testing.T) {
	out := t.TempDir()
	never().MkdirAll(out + "/o")
	w := &writer{out: out, fs: never(), o: small().withDefaults(), inum: map[string]uint32{}, dels: []string{"a"}}
	w.o.Codec = Stored{}
	if err := w.closeData(); err != nil || len(w.objs) != 1 {
		t.Fatal(err)
	}
	w.fs = &fsFuncs{create: 0}
	if err := w.closeData(); err == nil {
		t.Fatal("start")
	}
}
