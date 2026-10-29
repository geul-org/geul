//ff:func feature=pack type=resolver control=sequence
//ff:what TestReaderFindInLayer — descends internal blocks to the leaf; misses before the first child; bad child and cycles fail
package pack

import (
	"context"
	"errors"
	"testing"
)

func TestReaderFindInLayer(t *testing.T) {
	ctx := context.Background()
	_, out := built(t, small())
	r := opened(t, out)
	l := r.m.pLayers[0]
	if l.blk.kind != blkPInner {
		t.Fatal("fixture should have an internal P root")
	}
	if loc, ok, err := r.findInLayer(ctx, l, pKey(0, 0x845, 1)); err != nil || !ok || loc.Len == 0 {
		t.Fatal(loc, ok, err)
	}
	if _, ok, err := r.findInLayer(ctx, l, pKey(0, 0x846, 0)); err != nil || ok {
		t.Fatal("miss", err)
	}
	if _, ok, err := r.findInLayer(ctx, l, pKey(0, 1, 0)); err != nil || ok {
		t.Fatal("before first", err)
	}
	bad := l
	bad.blk = block{kind: blkPInner, keys: l.blk.keys, locs: append([]Loc{{Obj: 99}}, l.blk.locs[1:]...)}
	if _, _, err := r.findInLayer(ctx, bad, l.blk.keys[0]); !errors.Is(err, ErrCorrupt) {
		t.Fatal(err)
	}
	wrong := l
	wrong.blk = block{kind: blkPInner, keys: []string{pKey(0, 0, 0)}, locs: l.blk.locs[1:2]}
	if _, _, err := r.findInLayer(ctx, wrong, pKey(0, 1, 0)); !errors.Is(err, ErrCorrupt) {
		t.Fatal("child key", err)
	}
	self := []byte{blkPInner, 1, 0, 0, 0, 0, 65, 9, encRaw} // one child: itself (object 0, offset 64, 9 bytes, raw)
	obj, o := objBytes(kindIndex, self, nil)
	rr := newReader(memStore{objectName(o): obj}, Options{}, &manifest{objects: []object{o}}, "")
	cyc, _ := decodeBlock(self)
	if _, _, err := rr.findInLayer(ctx, layer{blk: cyc}, pKey(0, 0, 0)); !errors.Is(err, ErrCorrupt) {
		t.Fatal("cycle", err)
	}
}
