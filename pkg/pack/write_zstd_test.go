//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what TestWriteZstd — zstd with a trained dictionary: every path reads back, verify passes, the same input gives the same objects, records shrink; nodict compresses without one
package pack

import (
	"bytes"
	"context"
	"path/filepath"
	"testing"
)

func TestWriteZstd(t *testing.T) {
	ctx := context.Background()
	paths, get := synth(300)
	build := func(c Codec) (string, BuildResult) {
		out := filepath.Join(t.TempDir(), "z")
		o := small()
		o.Codec = c
		res, err := Write(ctx, out, Batch{Put: paths, Full: true, Get: get}, o)
		if err != nil {
			t.Fatal(err)
		}
		return out, res
	}
	out, res := build(&Zstd{DictSize: 4096, Workers: 3})
	r := opened(t, out)
	if res.Dict != FirstDictID || res.Compressed != len(paths) || len(r.m.dicts) != 1 || r.Verify(ctx) != nil {
		t.Fatal(res, r.m.dicts)
	}
	for _, p := range paths {
		want, _ := get(p)
		if b, err := r.Fetch(ctx, p); err != nil || !bytes.Equal(b, want) {
			t.Fatal(p, err)
		}
	}
	_, again := build(&Zstd{DictSize: 4096, Workers: 1})
	if again.Manifest != res.Manifest {
		t.Fatal("not deterministic across worker counts")
	}
	_, raw := build(Stored{})
	_, nodict := build(&Zstd{DictSize: -1})
	if nodict.Dict != 0 || nodict.Compressed != len(paths) || !(res.Bytes < nodict.Bytes && nodict.Bytes < raw.Bytes) {
		t.Fatal(res.Bytes, nodict.Bytes, raw.Bytes)
	}
}
