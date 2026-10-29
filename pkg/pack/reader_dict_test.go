//ff:func feature=pack type=resolver control=iteration dimension=1
//ff:what TestReaderDict — read once and cached; unknown ID, bad location and SHA-256 mismatch fail; a dictionary frame decodes through it
package pack

import (
	"context"
	"crypto/sha256"
	"errors"
	"testing"
)

func TestReaderDict(t *testing.T) {
	ctx := context.Background()
	obj, o := objBytes(kindData, []byte("DICT"), nil)
	m := &manifest{gen: 1, objects: []object{o}, dicts: []dictEntry{
		{id: 40000, loc: Loc{Off: 64, Len: 4}, hash: sha256.Sum256([]byte("DICT"))},
		{id: 40001, loc: Loc{Off: 64, Len: 4}},
		{id: 40002, loc: Loc{Obj: 5}},
	}}
	s := memStore{objectName(o): obj}
	r := newReader(s, Options{}, m, "")
	if d, err := r.dict(ctx, 40000); err != nil || string(d) != "DICT" {
		t.Fatal(err)
	}
	delete(s, objectName(o))
	if d, err := r.dict(ctx, 40000); err != nil || string(d) != "DICT" {
		t.Fatal("cache", err)
	}
	s[objectName(o)] = obj
	for _, id := range []uint32{1, 40001, 40002} {
		if _, err := r.dict(ctx, id); !errors.Is(err, ErrCorrupt) {
			t.Fatal(id, err)
		}
	}
	frame := []byte{0x28, 0xb5, 0x2f, 0xfd, 0x22, 0x40, 0x9c, 0x03, 0x19, 0x00, 0x00, 'a', 'b', 'c'}
	if b, err := r.unzstd(ctx, frame, false); err != nil || string(b) != "abc" {
		t.Fatal(err)
	}
}
