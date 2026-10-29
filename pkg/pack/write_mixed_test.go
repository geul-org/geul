//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what TestWriteMixed — a raw generation, then zstd, then one with no codec given (follows zstd, reuses the dictionary), then compaction (a new dictionary ID): every path of every layer reads back
package pack

import (
	"bytes"
	"context"
	"path/filepath"
	"testing"
)

func TestWriteMixed(t *testing.T) {
	ctx := context.Background()
	paths, get := synth(240)
	out := filepath.Join(t.TempDir(), "m")
	steps := []struct {
		put   []string
		codec Codec
	}{{paths[:80], nil}, {paths[80:160], &Zstd{DictSize: 4096}}, {paths[160:], nil}, {nil, &Zstd{DictSize: 4096}}}
	for i, s := range steps {
		o := small()
		o.Codec = s.codec
		var res BuildResult
		var err error
		if s.put == nil {
			res, err = Compact(ctx, out, o)
		} else {
			res, err = Write(ctx, out, Batch{Put: s.put, Get: get}, o)
		}
		want := [][2]int{{0, 0}, {80, FirstDictID}, {80, 0}, {240, FirstDictID + 1}}[i]
		if err != nil || res.Compressed != want[0] || res.Dict != uint32(want[1]) {
			t.Fatal(i, err, res.Compressed, res.Dict)
		}
	}
	r := opened(t, out)
	if r.Verify(ctx) != nil || len(r.m.dicts) != 1 || r.m.dicts[0].id != FirstDictID+1 {
		t.Fatal(r.m.dicts)
	}
	for _, p := range paths {
		want, _ := get(p)
		if b, err := r.Fetch(ctx, p); err != nil || !bytes.Equal(b, want) {
			t.Fatal(p, err)
		}
	}
}
