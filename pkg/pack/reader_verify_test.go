//ff:func feature=pack type=rule control=sequence
//ff:what TestReaderVerify — a built set passes; object damage, a dictionary SHA-256 mismatch, inventory/index disagreement and index errors fail
package pack

import (
	"context"
	"errors"
	"testing"
)

func TestReaderVerify(t *testing.T) {
	ctx := context.Background()
	_, out := built(t, small())
	r := opened(t, out)
	if err := r.Verify(ctx); err != nil {
		t.Fatal(err)
	}
	named := r.m.named
	r.m.named = named[:1]
	if err := r.Verify(ctx); !errors.Is(err, ErrCorrupt) {
		t.Fatal("disagree", err)
	}
	r.m.named = named
	r.m.pLayers = []layer{{blk: block{kind: blkPInner, keys: []string{"a"}, locs: []Loc{{Obj: 99}}}}}
	if err := r.Verify(ctx); !errors.Is(err, ErrCorrupt) {
		t.Fatal("index", err)
	}
	zout, _, _ := zstdSet(t, 100, &Zstd{DictSize: 4096})
	rz := opened(t, zout)
	if err := rz.Verify(ctx); err != nil {
		t.Fatal(err)
	}
	rz = opened(t, zout)
	rz.m.dicts[0].hash[0] ^= 1
	if err := rz.Verify(ctx); !errors.Is(err, ErrCorrupt) {
		t.Fatal("dictionary", err)
	}
	r2 := opened(t, out)
	r2.m.objects = append(r2.m.objects, object{kind: kindData, size: 64})
	if err := r2.Verify(ctx); err == nil {
		t.Fatal("objects")
	}
	obj, o := objBytes(kindData, nil, zstdFrame(encodeInventory(inventory{entries: []invEntry{{kind: 9}}})))
	r3 := newReader(memStore{objectName(o): obj}, Options{}, &manifest{objects: []object{o}}, "")
	if err := r3.Verify(ctx); !errors.Is(err, ErrCorrupt) {
		t.Fatal("state", err)
	}
}
