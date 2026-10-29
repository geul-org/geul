//ff:func feature=pack type=resolver control=iteration dimension=1
//ff:what TestReaderInventory — a data pack's inventory; unreadable or wrong headers and a broken inventory frame fail
package pack

import (
	"context"
	"errors"
	"testing"
)

func TestReaderInventory(t *testing.T) {
	ctx := context.Background()
	inv := sampleInventory()
	obj, o := objBytes(kindData, []byte("abc"), zstdFrame(encodeInventory(inv)))
	idx, io := objBytes(kindIndex, nil, nil)
	bad, bo := objBytes(kindData, nil, []byte{1, 2})
	s := memStore{objectName(o): obj, objectName(io): idx, objectName(bo): bad}
	r := newReader(s, Options{}, &manifest{objects: []object{o, io, bo, {kind: kindData, size: 64, hash: [32]byte{7}}}}, "")
	if got, err := r.inventory(ctx, 0); err != nil || len(got.entries) != 3 {
		t.Fatal(got, err)
	}
	for _, i := range []uint32{1, 2, 3} {
		if _, err := r.inventory(ctx, i); err == nil {
			t.Fatal(i)
		}
	}
	s[objectName(o)] = append([]byte{}, obj...)
	s[objectName(o)][0] = 'X'
	if _, err := r.inventory(ctx, 0); !errors.Is(err, ErrCorrupt) {
		t.Fatal(err)
	}
}
