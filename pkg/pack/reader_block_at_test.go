//ff:func feature=pack type=resolver control=iteration dimension=1
//ff:what TestReaderBlockAt — decoded and cached (cache dropped when full); bad location, bad frame and bad block fail
package pack

import (
	"context"
	"errors"
	"testing"
)

func TestReaderBlockAt(t *testing.T) {
	blk := zstdFrame(encodeBlock(block{kind: blkPLeaf, keys: []string{pKey(0, 1, 0)}, locs: []Loc{{Off: 64}}}))
	body := append(append(append([]byte{}, blk...), blk...), zstdFrame([]byte{9})...)
	obj, o := objBytes(kindIndex, append(body, 1, 2, 3), nil)
	s := memStore{objectName(o): obj}
	r, ctx := newReader(s, Options{CacheBlocks: 1}, &manifest{objects: []object{o}}, ""), context.Background()
	n := uint64(len(blk))
	if b, err := r.blockAt(ctx, Loc{Off: 64, Len: n, Enc: encZstd}); err != nil || b.kind != blkPLeaf {
		t.Fatal(err)
	}
	if _, err := r.blockAt(ctx, Loc{Off: 64, Len: n, Enc: encZstd}); err != nil {
		t.Fatal("cached", err)
	}
	if _, err := r.blockAt(ctx, Loc{Off: 64 + n, Len: n, Enc: encZstd}); err != nil || len(r.blocks) != 1 {
		t.Fatal("evict", err, len(r.blocks))
	}
	bad := []Loc{{Obj: 9}, {Off: 64, Len: 3, Enc: encZstd}, {Off: 64 + 2*n, Len: uint64(len(zstdFrame([]byte{9}))), Enc: encZstd}}
	for i, l := range bad {
		if _, err := r.blockAt(ctx, l); !errors.Is(err, ErrCorrupt) {
			t.Fatal(i, err)
		}
	}
}
