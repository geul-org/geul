//ff:func feature=pack type=decoder control=sequence
//ff:what TestReaderUnstore — raw as is, zstd unframed, other encodings ErrCorrupt
package pack

import (
	"context"
	"errors"
	"testing"
)

func TestReaderUnstore(t *testing.T) {
	r, ctx := newReader(memStore{}, Options{}, &manifest{}, ""), context.Background()
	if b, err := r.unstore(ctx, encRaw, []byte("x"), true); err != nil || string(b) != "x" {
		t.Fatal(err)
	}
	if b, err := r.unstore(ctx, encZstd, zstdFrame([]byte("y")), true); err != nil || string(b) != "y" {
		t.Fatal(err)
	}
	if _, err := r.unstore(ctx, encGone, nil, true); !errors.Is(err, ErrCorrupt) {
		t.Fatal(err)
	}
}
