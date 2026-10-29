//ff:func feature=pack type=decoder control=sequence
//ff:what TestReaderDecoder — 0 is the shared decoder; a DICTS dictionary loads once and is cached; missing or bad dictionaries fail
package pack

import (
	"context"
	"crypto/sha256"
	"errors"
	"os"
	"testing"
)

func TestReaderDecoder(t *testing.T) {
	ctx := context.Background()
	dict, _ := os.ReadFile("testdata/zstd/cli.dict")
	obj, o := objBytes(kindData, dict, nil)
	e := dictEntry{id: FirstDictID, loc: Loc{Off: headerSize, Len: uint64(len(dict))}, hash: sha256.Sum256(dict)}
	bad := dictEntry{id: FirstDictID + 1, loc: e.loc, hash: e.hash}
	r := newReader(memStore{objectName(o): obj}, Options{}, &manifest{objects: []object{o}, dicts: []dictEntry{e, bad}}, "")
	if d, err := r.decoder(ctx, 0); err != nil || d != plainDecoder() {
		t.Fatal(err)
	}
	d, err := r.decoder(ctx, FirstDictID)
	if err != nil {
		t.Fatal(err)
	}
	frame, _ := os.ReadFile("testdata/zstd/sample.txt.zst")
	if _, err := d.DecodeAll(frame, nil); err != nil {
		t.Fatal(err)
	}
	if again, _ := r.decoder(ctx, FirstDictID); again != d {
		t.Fatal("not cached")
	}
	if _, err := r.decoder(ctx, 9); !errors.Is(err, ErrCorrupt) {
		t.Fatal(err)
	}
	if _, err := r.decoder(ctx, FirstDictID+1); !errors.Is(err, ErrCorrupt) {
		t.Fatal(err)
	}
}
