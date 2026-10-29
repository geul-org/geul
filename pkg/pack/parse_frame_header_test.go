//ff:func feature=pack type=decoder control=iteration dimension=1
//ff:what TestParseFrameHeader — our frames, a zstd CLI frame (single segment, 1-byte size), 2-byte size (+256), dictionary IDs, bad magic, reserved bit, truncation
package pack

import (
	"errors"
	"testing"
)

func TestParseFrameHeader(t *testing.T) {
	h, err := parseFrameHeader(zstdFrame([]byte("abc")))
	if err != nil || h.size != 14 || h.fcs != 3 || !h.hasFCS || !h.checksum || h.window != 128<<10 {
		t.Fatal(h, err)
	}
	cli := []byte{0x28, 0xb5, 0x2f, 0xfd, 0x24, 0x03, 0x19, 0x00, 0x00, 'a', 'b', 'c', 0x99, 0x09, 0x77, 0xad}
	if h, err := parseFrameHeader(cli); err != nil || h.size != 6 || h.fcs != 3 || h.window != 3 {
		t.Fatal(h, err)
	}
	two := []byte{0x28, 0xb5, 0x2f, 0xfd, 0x61, 0x07, 0xe8, 0x02}
	if h, err := parseFrameHeader(two); err != nil || h.fcs != 1000 || h.dictID != 7 {
		t.Fatal(h, err)
	}
	for _, b := range [][]byte{{1, 2, 3, 4, 0}, {0x28, 0xb5, 0x2f, 0xfd, 0x08}, {0x28, 0xb5, 0x2f, 0xfd, 0xC4, 0x38, 1}} {
		if _, err := parseFrameHeader(b); !errors.Is(err, ErrCorrupt) {
			t.Fatal(b, err)
		}
	}
}
