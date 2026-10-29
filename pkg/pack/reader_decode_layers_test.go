//ff:func feature=pack type=decoder control=iteration dimension=1
//ff:what TestReaderDecodeLayers — roots decode in place; a bad frame, a bad block and the wrong key space fail
package pack

import (
	"context"
	"errors"
	"testing"
)

func TestReaderDecodeLayers(t *testing.T) {
	r, ctx := newReader(memStore{}, Options{}, &manifest{}, ""), context.Background()
	p := zstdFrame(encodeBlock(block{kind: blkPLeaf, keys: []string{pKey(0, 1, 0)}, locs: []Loc{{Off: 64}}}))
	ls := []layer{{enc: encZstd, root: p}}
	if err := r.decodeLayers(ctx, ls, true); err != nil || ls[0].blk.kind != blkPLeaf {
		t.Fatal(err)
	}
	for i, l := range []layer{{enc: encZstd, root: []byte{1}}, {enc: encRaw, root: []byte{9}}, {enc: encZstd, root: p}} {
		if err := r.decodeLayers(ctx, []layer{l}, i != 2); !errors.Is(err, ErrCorrupt) {
			t.Fatal(i, err)
		}
	}
}
