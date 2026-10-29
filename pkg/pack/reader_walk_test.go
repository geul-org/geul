//ff:func feature=pack type=resolver control=sequence
//ff:what TestReaderWalk — all leaf entries of a depth-3 H layer in order, a sub-range, callback errors and bad children stop it, depth guard
package pack

import (
	"context"
	"errors"
	"testing"
)

func TestReaderWalk(t *testing.T) {
	ctx := context.Background()
	_, out := built(t, small())
	r := opened(t, out)
	var keys []string
	collect := func(k string, _ Loc) error { keys = append(keys, k); return nil }
	if err := r.walk(ctx, r.m.hLayers[0].blk, "", "", 0, collect); err != nil || len(keys) != 8 || checkAscending(keys) != nil {
		t.Fatal(len(keys), err)
	}
	lo, hi := keys[2], keys[5]
	keys = nil
	if err := r.walk(ctx, r.m.hLayers[0].blk, lo, hi, 0, collect); err != nil || len(keys) != 3 || keys[0] != lo {
		t.Fatal(len(keys), err)
	}
	stop := errors.New("stop")
	if err := r.walk(ctx, r.m.hLayers[0].blk, "", "", 0, func(string, Loc) error { return stop }); err != stop {
		t.Fatal(err)
	}
	bad := block{kind: blkHInner, keys: []string{"a"}, locs: []Loc{{Obj: 99}}}
	if err := r.walk(ctx, bad, "", "", 0, collect); !errors.Is(err, ErrCorrupt) {
		t.Fatal(err)
	}
	if err := r.walk(ctx, bad, "", "", maxDepth, collect); !errors.Is(err, ErrCorrupt) {
		t.Fatal(err)
	}
}
