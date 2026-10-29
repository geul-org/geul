//ff:func feature=pack type=writer control=sequence
//ff:what TestWriterWriteData — records go into packs split at PackSize (an oversized record alone); a failing record stops it
package pack

import (
	"testing"
)

func TestWriterWriteData(t *testing.T) {
	out := t.TempDir()
	get := memGet(map[string][]byte{"a": make([]byte, 100), "b": make([]byte, 10), "c": make([]byte, 10)})
	w := testWriter(t, out, Batch{Put: []string{"a", "b", "c"}, Get: get})
	w.o.PackSize = 64 + 30
	if err := w.plan(); err != nil {
		t.Fatal(err)
	}
	never().MkdirAll(out + "/o")
	if err := w.writeData(); err != nil || len(w.objs) != 2 || w.res.Records != 3 {
		t.Fatal(err, len(w.objs))
	}
	w = testWriter(t, out, Batch{Put: []string{"a"}, Get: get})
	w.plan()
	w.b.Get = memGet(nil)
	if err := w.writeData(); err == nil {
		t.Fatal("get")
	}
}
