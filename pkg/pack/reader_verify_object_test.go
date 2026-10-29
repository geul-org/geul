//ff:func feature=pack type=rule control=iteration dimension=1
//ff:what TestReaderVerifyObject — a good object passes; missing, bad header, wrong kind, wrong hash, short and over-long objects fail
package pack

import (
	"context"
	"crypto/sha256"
	"errors"
	"testing"
)

func TestReaderVerifyObject(t *testing.T) {
	ctx := context.Background()
	obj, o := objBytes(kindIndex, []byte("abc"), nil)
	junk := append([]byte{}, obj...)
	junk[0] = 'X'
	s := memStore{objectName(o): obj}
	objs := []object{o, {kind: kindIndex, size: 67, hash: [32]byte{1}}, {kind: kindData, size: o.size, hash: o.hash}, {kind: kindIndex, size: o.size, hash: [32]byte{2}}, {kind: kindIndex, size: o.size, hash: [32]byte{3}}}
	s[objectName(objs[2])] = obj
	s[objectName(objs[3])] = obj
	s[objectName(objs[4])] = junk
	r := newReader(s, Options{}, &manifest{objects: objs}, "")
	if err := r.verifyObject(ctx, 0); err != nil {
		t.Fatal(err)
	}
	if err := r.verifyObject(ctx, 1); err == nil {
		t.Fatal("missing")
	}
	for _, i := range []int{2, 3, 4} {
		if err := r.verifyObject(ctx, i); !errors.Is(err, ErrCorrupt) {
			t.Fatal(i, err)
		}
	}
	short, so := objBytes(kindIndex, []byte("abcdef"), nil)
	so.size = 64 + 3
	copy(short[16:24], []byte{67, 0, 0, 0, 0, 0, 0, 0})
	so.hash = sha256.Sum256(short[:67])
	r = newReader(memStore{objectName(so): short}, Options{}, &manifest{objects: []object{so}}, "")
	if err := r.verifyObject(ctx, 0); !errors.Is(err, ErrCorrupt) {
		t.Fatal("long", err)
	}
	trunc := append([]byte{}, obj[:65]...)
	r = newReader(memStore{objectName(o): trunc}, Options{}, &manifest{objects: []object{o}}, "")
	if err := r.verifyObject(ctx, 0); !errors.Is(err, ErrCorrupt) {
		t.Fatal("short", err)
	}
}
