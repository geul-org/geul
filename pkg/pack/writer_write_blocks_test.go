//ff:func feature=pack type=writer control=sequence
//ff:what TestWriterWriteBlocks — checksummed frames in index objects, a new object when full; codec, close and start failures
package pack

import (
	"errors"
	"testing"
)

func TestWriterWriteBlocks(t *testing.T) {
	out := t.TempDir()
	never().MkdirAll(out + "/o")
	bs := chunk([]kv{{"a", Loc{Off: 64}}, {"b", Loc{Off: 65}}}, blkHLeaf, 1)
	w := &writer{out: out, fs: never(), o: BuildOptions{PackSize: 64 + 30}.withDefaults()}
	locs, err := w.writeBlocks(bs)
	if err != nil || locs[1].Obj != 1 || locs[0].Enc != encZstd || len(w.objs) != 1 {
		t.Fatal(locs, err)
	}
	w.o.Codec = fakeCodec{err: errors.New("x")}
	if _, err := w.writeBlocks(bs); err == nil {
		t.Fatal("codec")
	}
	w.o.Codec = Stored{}
	w.fs = &fsFuncs{mkdir: -1, create: -1, rename: 0}
	if _, err := w.writeBlocks(bs); err == nil {
		t.Fatal("close")
	}
	w.fs = &fsFuncs{create: 0}
	if _, err := w.writeBlocks(bs); err == nil {
		t.Fatal("start")
	}
}
