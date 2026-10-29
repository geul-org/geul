//ff:func feature=pack type=writer control=sequence
//ff:what TestStreamLayerOf — one zstd layer with the count and the root block; a failing codec fails
package pack

import "testing"

func TestStreamLayerOf(t *testing.T) {
	ls, err := streamTestNew(t, nil).layerOf(blockOf(blkPLeaf, []streamKV{{Key: pKey(0, 1, 0)}}), 1)
	if err != nil || len(ls) != 1 || ls[0].count != 1 || ls[0].enc != encZstd {
		t.Fatal(err)
	}
	if b, err := decodeBlock(mustUnzstd(t, ls[0].root)); err != nil || b.kind != blkPLeaf {
		t.Fatal(err)
	}
	if _, err := streamTestNew(t, failCodec{}).layerOf(block{kind: blkPLeaf}, 1); err == nil {
		t.Fatal("failing codec: want error")
	}
}
