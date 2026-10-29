//ff:func feature=pack type=decoder control=iteration dimension=1
//ff:what TestReaderUnzstd — content size required, checksum when asked, dictionary looked up, codec errors and size mismatch fail
package pack

import (
	"context"
	"errors"
	"testing"
)

func TestReaderUnzstd(t *testing.T) {
	r, ctx := newReader(memStore{}, Options{}, &manifest{}, ""), context.Background()
	noFCS := []byte{0x28, 0xb5, 0x2f, 0xfd, 0x04, 0x38, 0x01, 0, 0, 0, 0, 0, 0}
	noSum := []byte{0x28, 0xb5, 0x2f, 0xfd, 0x20, 0x03, 0x19, 0x00, 0x00, 'a', 'b', 'c'}
	if b, err := r.unzstd(ctx, noSum, false); err != nil || string(b) != "abc" {
		t.Fatal(err)
	}
	dict := []byte{0x28, 0xb5, 0x2f, 0xfd, 0x21, 0x07, 0x03, 0x19, 0x00, 0x00, 'a', 'b', 'c'}
	lie := []byte{0x28, 0xb5, 0x2f, 0xfd, 0x20, 0x04, 0x19, 0x00, 0x00, 'a', 'b', 'c'}
	for i, f := range [][]byte{noFCS, noSum, dict, lie, {1}} {
		if _, err := r.unzstd(ctx, f, i == 1); !errors.Is(err, ErrCorrupt) {
			t.Fatal(i, err)
		}
	}
	rep := []byte{0x28, 0xb5, 0x2f, 0xfd, 0x64, 0xe8, 0x02, 0x45, 0x00, 0x00, 0x08, 0x78, 0x01, 0x00, 0xe4, 0x2b, 0x20, 0x04, 0xe1, 0x00, 0xb7, 0x9c}
	if _, err := r.unzstd(ctx, rep, true); !errors.Is(err, ErrUnsupported) {
		t.Fatal(err)
	}
}
