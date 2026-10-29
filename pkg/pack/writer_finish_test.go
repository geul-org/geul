//ff:func feature=pack type=writer control=sequence
//ff:what TestWriterFinish — manifest chained to the previous one, latest.txt last; layer, manifest, verify, conflict and latest failures
package pack

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestWriterFinish(t *testing.T) {
	_, out := built(t, small())
	prevName := opened(t, out).ManifestName()
	mk := func() *writer {
		w := testWriter(t, out, Batch{Put: []string{"z.geul"}, Get: memGet(map[string][]byte{"z.geul": {1}})})
		w.plan()
		w.writeData()
		return w
	}
	w := mk()
	w.o.Codec = fakeCodec{err: errors.New("x")}
	if err := w.finish(); err == nil {
		t.Fatal("layers")
	}
	w = mk()
	w.fs = &fsFuncs{create: 1, mkdir: -1, rename: -1}
	if err := w.finish(); err == nil {
		t.Fatal("manifest")
	}
	w = mk()
	w.o.Verify = false
	w.o.Check = func(_ context.Context, _ *Reader) error { return errors.New("check") }
	if err := w.finish(); err == nil {
		t.Fatal("check")
	}
	w = mk()
	os.WriteFile(filepath.Join(out, latestFile), []byte("9 x\n"), 0o644)
	if err := w.finish(); !errors.Is(err, ErrConflict) {
		t.Fatal(err)
	}
	os.WriteFile(filepath.Join(out, latestFile), []byte("1 "+prevName+"\n"), 0o644)
	w = mk()
	w.fs = &fsFuncs{create: 2, mkdir: -1, rename: -1}
	if err := w.finish(); err == nil {
		t.Fatal("latest")
	}
	w = mk()
	if err := w.finish(); err != nil || w.res.Generation != 2 {
		t.Fatal(err)
	}
	r := opened(t, out)
	if prev, _ := parseHashHex(prevName); r.m.prev != prev || r.Generation() != 2 {
		t.Fatal("chain")
	}
}
