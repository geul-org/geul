//ff:func feature=pack type=decoder control=iteration dimension=1
//ff:what TestReaderUnzstd — raw and compressed (zstd CLI) frames decode; missing content size or checksum, bad blocks, trailing bytes, unknown dictionary, size and checksum mismatches are ErrCorrupt
package pack

import (
	"bytes"
	"context"
	"errors"
	"testing"
)

func TestReaderUnzstd(t *testing.T) {
	r, ctx := newReader(memStore{}, Options{}, &manifest{}, ""), context.Background()
	noSum := []byte{0x28, 0xb5, 0x2f, 0xfd, 0x20, 0x03, 0x19, 0x00, 0x00, 'a', 'b', 'c'}
	if b, err := r.unzstd(ctx, noSum, false); err != nil || string(b) != "abc" {
		t.Fatal(err)
	}
	// `printf x%.0s {1..1000} | zstd -c`: one compressed block
	rep := []byte{0x28, 0xb5, 0x2f, 0xfd, 0x64, 0xe8, 0x02, 0x45, 0x00, 0x00, 0x08, 0x78, 0x01, 0x00, 0xe4, 0x2b, 0x20, 0x04, 0xe1, 0x00, 0xb7, 0x9c}
	if b, err := r.unzstd(ctx, rep, true); err != nil || !bytes.Equal(b, bytes.Repeat([]byte("x"), 1000)) {
		t.Fatal(err)
	}
	noFCS := []byte{0x28, 0xb5, 0x2f, 0xfd, 0x04, 0x38, 0x01, 0, 0, 0, 0, 0, 0}
	truncated := []byte{0x28, 0xb5, 0x2f, 0xfd, 0x20, 0x03, 0x19, 0x00}
	trailing := append(append([]byte{}, noSum...), 0)
	dict := []byte{0x28, 0xb5, 0x2f, 0xfd, 0x21, 0x07, 0x03, 0x19, 0x00, 0x00, 'a', 'b', 'c'}
	lie := []byte{0x28, 0xb5, 0x2f, 0xfd, 0x20, 0x04, 0x19, 0x00, 0x00, 'a', 'b', 'c'}
	badSum := append(append([]byte{}, rep[:len(rep)-1]...), 0)
	for i, f := range [][]byte{noFCS, noSum, truncated, trailing, dict, lie, badSum, {1}} {
		if _, err := r.unzstd(ctx, f, i == 1); !errors.Is(err, ErrCorrupt) {
			t.Fatal(i, err)
		}
	}
}
