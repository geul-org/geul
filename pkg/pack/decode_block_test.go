//ff:func feature=pack type=decoder control=iteration dimension=1
//ff:what TestDecodeBlock — header, P and H columns, trailing bytes and key order are checked
package pack

import (
	"errors"
	"strings"
	"testing"
)

func TestDecodeBlock(t *testing.T) {
	ok := encodeBlock(block{kind: blkPLeaf, keys: []string{pKey(0, 1, 0)}, locs: []Loc{{Off: 64, Len: 1}}})
	if _, err := decodeBlock(ok); err != nil {
		t.Fatal(err)
	}
	h := strings.Repeat("b", 32) + strings.Repeat("a", 32)
	bad := [][]byte{
		{}, {5, 0}, {1, 100},
		{1, 1, 0x80},
		{1, 1, 0, 1, 0, 0, 65, 1, 7},
		append(ok, 0),
		append(append([]byte{3, 2}, h...), 0, 0, 65, 66, 1, 1, 0, 0),
	}
	for i, b := range bad {
		if _, err := decodeBlock(b); !errors.Is(err, ErrCorrupt) {
			t.Fatal(i, err)
		}
	}
}
