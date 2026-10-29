//ff:func feature=pack type=writer control=sequence
//ff:what TestWriterOpen — empty out, a previous generation (OBJECTS grows), Fresh, the H layer limit (OBJECTS rebuilt), and errors
package pack

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestWriterOpen(t *testing.T) {
	ctx := context.Background()
	w := &writer{ctx: ctx, out: t.TempDir(), o: small().withDefaults()}
	if err := w.open(); err != nil || w.prev != nil || !w.all {
		t.Fatal(err)
	}
	_, out := built(t, small())
	w = &writer{ctx: ctx, out: out, o: small().withDefaults()}
	if err := w.open(); err != nil || w.all || len(w.base) != 18 || len(w.objs) != 2 || w.res.Generation != 1 {
		t.Fatal(err, w.all, len(w.base))
	}
	w = &writer{ctx: ctx, out: out, o: BuildOptions{Fresh: true}.withDefaults()}
	if err := w.open(); err != nil || len(w.base) != 0 || len(w.issuers) != 1 || w.objs != nil {
		t.Fatal(err)
	}
	w = &writer{ctx: ctx, out: out, o: BuildOptions{HLayers: 1}.withDefaults()}
	if err := w.open(); err != nil || !w.all || len(w.objs) != 1 || w.remap == nil {
		t.Fatal(err, w.all)
	}
	r := opened(t, out)
	flip(t, filepath.Join(out, "m", r.ManifestName()+".gpm"), 0)
	if err := (&writer{ctx: ctx, out: out, o: small().withDefaults()}).open(); !errors.Is(err, ErrCorrupt) {
		t.Fatal(err)
	}
	_, out2 := built(t, small())
	r2 := opened(t, out2)
	gpx := objPath(t, r2, out2, kindIndex)
	os.Remove(gpx)
	if err := (&writer{ctx: ctx, out: out2, o: small().withDefaults()}).open(); err == nil {
		t.Fatal("entries")
	}
}
