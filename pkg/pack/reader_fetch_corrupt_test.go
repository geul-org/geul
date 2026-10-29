//ff:func feature=pack type=resolver control=sequence
//ff:what TestReaderFetchCorrupt — a damaged zstd record fails its checksum and a damaged dictionary its SHA-256: ErrCorrupt, never wrong bytes
package pack

import (
	"context"
	"errors"
	"path/filepath"
	"testing"
)

func TestReaderFetchCorrupt(t *testing.T) {
	ctx := context.Background()
	paths, get := synth(100)
	out := filepath.Join(t.TempDir(), "c")
	o := small()
	o.Codec = &Zstd{DictSize: 4096}
	if _, err := Write(ctx, out, Batch{Put: paths, Get: get}, o); err != nil {
		t.Fatal(err)
	}
	r := opened(t, out)
	l := r.m.named[len(r.m.named)-1].loc
	gpk := objPath(t, r, out, kindData)
	flip(t, gpk, int(l.Off+l.Len)-6)
	if _, err := opened(t, out).Fetch(ctx, r.m.named[len(r.m.named)-1].path); !errors.Is(err, ErrCorrupt) {
		t.Fatal("record", err)
	}
	flip(t, gpk, int(r.m.dicts[0].loc.Off)+200)
	r2 := opened(t, out)
	if _, err := r2.Fetch(ctx, paths[0]); !errors.Is(err, ErrCorrupt) {
		t.Fatal("dictionary", err)
	}
	if err := r2.Verify(ctx); err == nil {
		t.Fatal("verify")
	}
}
