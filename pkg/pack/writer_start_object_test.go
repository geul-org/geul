//ff:func feature=pack type=writer control=sequence
//ff:what TestWriterStartObject — header placeholder, number = next OBJECTS place; a temp file failure is returned
package pack

import "testing"

func TestWriterStartObject(t *testing.T) {
	out := t.TempDir()
	never().MkdirAll(out + "/o")
	w := &writer{out: out, fs: never(), objs: []object{{}}}
	if err := w.startObject(kindIndex); err != nil || w.cur.off != 64 || w.curNum != 1 {
		t.Fatal(err)
	}
	w.abort()
	w.fs = &fsFuncs{create: 0}
	if err := w.startObject(kindIndex); err == nil {
		t.Fatal("create")
	}
}
