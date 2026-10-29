//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what TestWriterPrepareDict — none for Stored, nodict and too few records; a set's dictionary is reused; one is trained when it pays, not when it does not or training fails; sample and reuse errors are returned
package pack

import (
	"errors"
	"os"
	"strings"
	"testing"
)

func TestWriterPrepareDict(t *testing.T) {
	prep := func(out string, n int, c Codec, get func(string) ([]byte, error)) (*writer, error) {
		paths, sg := synth(n)
		if get == nil {
			get = sg
		}
		w := testWriter(t, out, Batch{Put: paths, Get: sg})
		w.o.Codec = c
		w.plan()
		w.b.Get = get
		never().MkdirAll(out + "/o")
		return w, w.prepareDict()
	}
	for _, c := range []Codec{Stored{}, &Zstd{DictSize: -1}} {
		if w, err := prep(t.TempDir(), 100, c, nil); err != nil || w.dict != nil {
			t.Fatal(c, err)
		}
	}
	if w, err := prep(t.TempDir(), 10, &Zstd{DictSize: 4096}, nil); err != nil || w.dict != nil {
		t.Fatal("few", err)
	}
	if w, err := prep(t.TempDir(), 100, &Zstd{DictSize: 4096}, nil); err != nil || w.dict == nil || w.res.Dict != FirstDictID || len(w.dicts) != 1 {
		t.Fatal("trained", err)
	}
	if w, err := prep(t.TempDir(), 100, &Zstd{}, nil); err != nil || w.dict != nil {
		t.Fatal("110 KiB does not pay on 100 small records", err)
	}
	same := func(string) ([]byte, error) { return []byte(strings.Repeat("x", 300)), nil }
	if w, err := prep(t.TempDir(), 100, &Zstd{DictSize: 4096}, same); err != nil || w.dict != nil {
		t.Fatal("training fails on uniform data", err)
	}
	if _, err := prep(t.TempDir(), 100, &Zstd{DictSize: 4096}, memGet(nil)); err == nil {
		t.Fatal("samples")
	}
	out, _, _ := zstdSet(t, 100, &Zstd{DictSize: 4096})
	w, err := prep(out, 120, &Zstd{DictSize: 4096}, nil)
	if err != nil || w.dict == nil || w.dict.ID != FirstDictID || w.res.Dict != 0 {
		t.Fatal("reuse", err)
	}
	os.Remove(objPath(t, w.prev, out, kindData))
	if _, err := prep(out, 120, &Zstd{DictSize: 4096}, nil); !errors.Is(err, ErrCorrupt) {
		t.Fatal("reuse read", err)
	}
}
