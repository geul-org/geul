//ff:func feature=pack type=decoder control=sequence
//ff:what TestZstdBlocks — blocks until Last_Block; a bad block stops with its error
package pack

import (
	"errors"
	"testing"
)

func TestZstdBlocks(t *testing.T) {
	out, rest, err := zstdBlocks([]byte{0x10, 0, 0, 'a', 'b', 0x09, 0, 0, 'c', 7}, 1<<20)
	if err != nil || string(out) != "abc" || len(rest) != 1 {
		t.Fatal(string(out), err)
	}
	if _, _, err := zstdBlocks([]byte{0x10, 0, 0, 'a'}, 1<<20); !errors.Is(err, ErrCorrupt) {
		t.Fatal(err)
	}
}
