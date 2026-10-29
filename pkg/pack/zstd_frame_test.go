//ff:func feature=pack type=encoder control=sequence
//ff:what TestZstdFrame — header bytes, raw blocks split at 128 KiB, empty input, and a round trip through Stored
package pack

import (
	"bytes"
	"testing"
)

func TestZstdFrame(t *testing.T) {
	f := zstdFrame(nil)
	if !bytes.Equal(f[:6], []byte{0x28, 0xB5, 0x2F, 0xFD, 0xC4, 0x38}) || len(f) != 6+8+3+4 {
		t.Fatalf("%x", f)
	}
	big := bytes.Repeat([]byte{1, 2, 3}, 100000)
	f = zstdFrame(big)
	if len(f) != 14+3*3+len(big)+4 {
		t.Fatal(len(f))
	}
	got, err := Stored{}.Decompress(f, nil)
	if err != nil || !bytes.Equal(got, big) {
		t.Fatal(err)
	}
}
