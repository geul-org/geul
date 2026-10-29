//ff:func feature=pack type=decoder control=iteration dimension=1
//ff:what TestZstdBlock — raw and RLE decode; compressed is ErrUnsupported; reserved, oversize and truncated blocks are ErrCorrupt
package pack

import (
	"errors"
	"testing"
)

func TestZstdBlock(t *testing.T) {
	out, rest, last, err := zstdBlock(nil, []byte{0x19, 0, 0, 'a', 'b', 'c', 9}, 100)
	if err != nil || string(out) != "abc" || !last || len(rest) != 1 {
		t.Fatal(out, err)
	}
	out, _, last, err = zstdBlock([]byte("x"), []byte{0x22, 0, 0, 'y'}, 100)
	if err != nil || string(out) != "xyyyy" || last {
		t.Fatal(string(out), err)
	}
	if _, _, _, err := zstdBlock(nil, []byte{0x45, 0, 0, 0}, 100); !errors.Is(err, ErrUnsupported) {
		t.Fatal(err)
	}
	for _, b := range [][]byte{{0x07, 0, 0}, {0x19, 0, 0, 'a'}, {0x22, 0, 0}, {0x19, 0}, {0x00, 0x08, 0}} {
		if _, _, _, err := zstdBlock(nil, b, 100); !errors.Is(err, ErrCorrupt) {
			t.Fatalf("%x %v", b, err)
		}
	}
}
