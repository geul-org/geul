//ff:func feature=pack type=writer control=sequence
//ff:what TestWriterAbort — removes an open object's temp file; nothing open is a no-op
package pack

import (
	"os"
	"testing"
)

func TestWriterAbort(t *testing.T) {
	out := t.TempDir()
	never().MkdirAll(out + "/o")
	w := &writer{out: out, fs: never()}
	w.abort()
	w.startObject(kindIndex)
	name := w.cur.f.Name()
	w.abort()
	if _, err := os.Stat(name); err == nil || w.cur != nil {
		t.Fatal("left")
	}
}
