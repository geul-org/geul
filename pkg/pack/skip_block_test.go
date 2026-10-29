//ff:func feature=pack type=decoder control=iteration dimension=1
//ff:what TestSkipBlock — raw, RLE and compressed blocks are stepped over; truncated, reserved and oversize blocks are ErrCorrupt
package pack

import (
	"errors"
	"testing"
)

func TestSkipBlock(t *testing.T) {
	rest, last, err := skipBlock([]byte{0x19, 0, 0, 'a', 'b', 'c', 9}, 100)
	if err != nil || !last || len(rest) != 1 || rest[0] != 9 {
		t.Fatal(rest, err)
	}
	if rest, last, err = skipBlock([]byte{0x22, 0, 0, 'y', 8}, 100); err != nil || last || len(rest) != 1 {
		t.Fatal(rest, err)
	}
	if rest, _, err = skipBlock([]byte{0x15, 0, 0, 1, 2, 7}, 100); err != nil || len(rest) != 1 {
		t.Fatal(rest, err)
	}
	for _, b := range [][]byte{{0x19, 0}, {0x07, 0, 0}, {0x19, 0, 0, 'a'}, {0x22, 0, 0}, {0x00, 0x08, 0}} {
		if _, _, err := skipBlock(b, 100); !errors.Is(err, ErrCorrupt) {
			t.Fatalf("%x %v", b, err)
		}
	}
}
