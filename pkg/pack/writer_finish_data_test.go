//ff:func feature=pack type=writer control=sequence
//ff:what TestWriterFinishData — the inventory frame closes the pack; codec errors fail
package pack

import (
	"errors"
	"testing"
)

func TestWriterFinishData(t *testing.T) {
	out := t.TempDir()
	never().MkdirAll(out + "/o")
	w := &writer{out: out, fs: never(), o: small().withDefaults()}
	w.startData()
	if err := w.finishData(); err != nil || len(w.objs) != 1 {
		t.Fatal(err)
	}
	w.startData()
	w.o.Codec = fakeCodec{err: errors.New("x")}
	if err := w.finishData(); err == nil {
		t.Fatal("codec")
	}
}
