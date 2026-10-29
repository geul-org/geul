//ff:func feature=pack type=decoder control=sequence
//ff:what TestFrameEnd — multi-block frames end after the checksum; no checksum, bad blocks and a cut checksum
package pack

import (
	"bytes"
	"errors"
	"testing"
)

func TestFrameEnd(t *testing.T) {
	f := zstdFrame(bytes.Repeat([]byte{7}, 300000))
	h, _ := parseFrameHeader(f)
	if n, err := frameEnd(append(f, 1, 2), h); err != nil || n != len(f) {
		t.Fatal(n, err)
	}
	noSum := []byte{0x28, 0xb5, 0x2f, 0xfd, 0x20, 0x03, 0x19, 0x00, 0x00, 'a', 'b', 'c', 5}
	h, _ = parseFrameHeader(noSum)
	if n, err := frameEnd(noSum, h); err != nil || n != 12 {
		t.Fatal(n, err)
	}
	h, _ = parseFrameHeader(f)
	if _, err := frameEnd(f[:20], h); !errors.Is(err, ErrCorrupt) {
		t.Fatal(err)
	}
	if _, err := frameEnd(f[:len(f)-2], h); !errors.Is(err, ErrCorrupt) {
		t.Fatal(err)
	}
}
