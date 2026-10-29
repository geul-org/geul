//ff:func feature=pack type=resolver control=sequence
//ff:what TestReaderEntries — P, H and named paths with locations; bad layers and out-of-range Issuer numbers fail
package pack

import (
	"context"
	"errors"
	"testing"
)

func TestReaderEntries(t *testing.T) {
	ctx := context.Background()
	_, out := built(t, small())
	r := opened(t, out)
	es, err := r.entries(ctx)
	if err != nil || len(es) != 18 {
		t.Fatal(len(es), err)
	}
	bad := []layer{{blk: block{kind: blkHInner, keys: []string{"a"}, locs: []Loc{{Obj: 99}}}}}
	h := r.m.hLayers
	r.m.hLayers = bad
	if _, err := r.entries(ctx); !errors.Is(err, ErrCorrupt) {
		t.Fatal(err)
	}
	r.m.hLayers = h
	r.m.pLayers = []layer{{blk: block{kind: blkPLeaf, keys: []string{pKey(5, 1, 0)}, locs: []Loc{{Off: 64}}}}}
	if _, err := r.entries(ctx); !errors.Is(err, ErrCorrupt) {
		t.Fatal(err)
	}
	r.m.pLayers = bad
	if _, err := r.entries(ctx); !errors.Is(err, ErrCorrupt) {
		t.Fatal(err)
	}
}
