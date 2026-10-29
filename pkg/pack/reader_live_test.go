//ff:func feature=pack type=resolver control=sequence
//ff:what TestReaderLive — newest layer wins, deletion markers drop keys, range limits apply, walk errors pass up
package pack

import (
	"context"
	"errors"
	"testing"
)

func TestReaderLive(t *testing.T) {
	r, ctx := newReader(memStore{}, Options{}, &manifest{}, ""), context.Background()
	newer := layer{blk: block{kind: blkHLeaf, keys: []string{"a", "b"}, locs: []Loc{{Enc: encGone}, {Off: 2}}}}
	older := layer{blk: block{kind: blkHLeaf, keys: []string{"a", "b", "c"}, locs: []Loc{{Off: 1}, {Off: 3}, {Off: 4}}}}
	s, err := r.live(ctx, []layer{newer, older}, "", "")
	if err != nil || len(s) != 2 || s["b"].Off != 2 || s["c"].Off != 4 {
		t.Fatal(s, err)
	}
	if s, _ := r.live(ctx, []layer{older}, "b", "c"); len(s) != 1 {
		t.Fatal(s)
	}
	bad := layer{blk: block{kind: blkHInner, keys: []string{"a"}, locs: []Loc{{Obj: 9}}}}
	if _, err := r.live(ctx, []layer{bad}, "", ""); !errors.Is(err, ErrCorrupt) {
		t.Fatal(err)
	}
}
