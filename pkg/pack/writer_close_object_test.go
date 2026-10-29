//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what TestWriterCloseObject — a closed object is named by its SHA-256 and parses; each file failure removes the temp file and fails
package pack

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
	"path/filepath"
	"testing"
)

func TestWriterCloseObject(t *testing.T) {
	out := t.TempDir()
	never().MkdirAll(out + "/o")
	w := &writer{out: out, fs: never()}
	w.startObject(kindData)
	w.cur.add([]byte("rec"))
	if err := w.closeObject([]byte("inv")); err != nil {
		t.Fatal(err)
	}
	b, _ := os.ReadFile(filepath.Join(out, w.res.Objects[0]))
	sum := sha256.Sum256(b)
	h, err := parseObjectHeader(b[:64])
	if err != nil || "o/"+hex.EncodeToString(sum[:])+".gpk" != w.res.Objects[0] || h.invOff != 67 || h.count != 1 {
		t.Fatal(err, h)
	}
	for i, bf := range []badFile{{write: true}, {writeAt: true}, {readAt: true}, {close: true}} {
		bf := bf
		w.fs = &fsFuncs{mkdir: -1, create: -1, rename: -1, wrap: func(f objFile) objFile { bf.objFile = f; return &bf }}
		w.startObject(kindIndex)
		if err := w.closeObject(nil); err == nil {
			t.Fatal(i)
		}
		if _, err := os.Stat(bf.objFile.Name()); err == nil {
			t.Fatal("temp left", i)
		}
	}
	w.fs = &fsFuncs{mkdir: -1, create: -1, rename: 0}
	w.startObject(kindIndex)
	if err := w.closeObject(nil); err == nil {
		t.Fatal("rename")
	}
}
