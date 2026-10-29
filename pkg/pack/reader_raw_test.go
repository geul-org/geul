//ff:func feature=pack type=resolver control=iteration dimension=1
//ff:what TestReaderRaw — one range inside the object; bad object number, wrong kind and out-of-bounds locations fail
package pack

import (
	"context"
	"errors"
	"testing"
)

func TestReaderRaw(t *testing.T) {
	obj, o := objBytes(kindData, []byte("abcd"), nil)
	r, ctx := newReader(memStore{objectName(o): obj}, Options{}, &manifest{objects: []object{o}}, ""), context.Background()
	if b, err := r.raw(ctx, kindData, Loc{Off: 65, Len: 2}); err != nil || string(b) != "bc" {
		t.Fatal(err)
	}
	for i, c := range []struct {
		kind byte
		l    Loc
	}{{kindData, Loc{Obj: 1}}, {kindIndex, Loc{Off: 64}}, {kindData, Loc{Off: 10}}, {kindData, Loc{Off: 99}}, {kindData, Loc{Off: 64, Len: 5}}} {
		if _, err := r.raw(ctx, c.kind, c.l); !errors.Is(err, ErrCorrupt) {
			t.Fatal(i, err)
		}
	}
}
