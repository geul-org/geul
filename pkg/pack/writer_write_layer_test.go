//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what TestWriterWriteLayer — no entries no layer; one leaf is the root; rootMax 0 gives depth 2; fanout levels give depth 3; failures
package pack

import (
	"errors"
	"strings"
	"testing"
)

func TestWriterWriteLayer(t *testing.T) {
	out := t.TempDir()
	never().MkdirAll(out + "/o")
	w := &writer{out: out, fs: never(), o: BuildOptions{}.withDefaults()}
	var kvs []kv
	for _, c := range "abcde" {
		kvs = append(kvs, kv{strings.Repeat(string(c), 32), Loc{Off: 64}})
	}
	if ls, err := w.writeLayer(nil, blkHLeaf, 2, 2, 2); ls != nil || err != nil {
		t.Fatal(ls, err)
	}
	ls, err := w.writeLayer(kvs, blkHLeaf, 8, 2, 0)
	r := newReader(memStore{}, Options{}, &manifest{}, "")
	if err != nil || r.decodeLayers(nil, ls, false) != nil || ls[0].blk.kind != blkHLeaf || ls[0].count != 5 {
		t.Fatal(err)
	}
	ls, _ = w.writeLayer(kvs, blkHLeaf, 2, 2, 0)
	r.decodeLayers(nil, ls, false)
	if ls[0].blk.kind != blkHInner || len(ls[0].blk.keys) != 3 {
		t.Fatal(ls[0].blk)
	}
	ls, _ = w.writeLayer(kvs, blkHLeaf, 1, 2, 2)
	r.decodeLayers(nil, ls, false)
	if len(ls[0].blk.keys) != 2 {
		t.Fatal(ls[0].blk)
	}
	w.o.Codec = fakeCodec{err: errors.New("x")}
	if _, err := w.writeLayer(kvs, blkHLeaf, 2, 2, 0); err == nil {
		t.Fatal("blocks")
	}
	if _, err := w.writeLayer(kvs[:1], blkHLeaf, 2, 2, 0); err == nil {
		t.Fatal("root")
	}
}
