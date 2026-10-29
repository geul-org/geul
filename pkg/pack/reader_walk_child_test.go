//ff:func feature=pack type=resolver control=sequence
//ff:what TestReaderWalkChild — walks one child; an unreadable child or one not matching its parent fails
package pack

import (
	"context"
	"errors"
	"testing"
)

func TestReaderWalkChild(t *testing.T) {
	ctx := context.Background()
	_, out := built(t, small())
	r := opened(t, out)
	root := r.m.pLayers[0].blk
	n := 0
	if err := r.walkChild(ctx, root, 0, "", "", 0, func(string, Loc) error { n++; return nil }); err != nil || n != 2 {
		t.Fatal(n, err)
	}
	wrong := block{kind: blkPInner, keys: []string{pKey(0, 0, 0)}, locs: root.locs[:1]}
	if err := r.walkChild(ctx, wrong, 0, "", "", 0, nil); !errors.Is(err, ErrCorrupt) {
		t.Fatal(err)
	}
	if err := r.walkChild(ctx, block{kind: blkPInner, keys: root.keys[:1], locs: []Loc{{Obj: 99}}}, 0, "", "", 0, nil); !errors.Is(err, ErrCorrupt) {
		t.Fatal(err)
	}
}
