//ff:func feature=pack type=decoder control=iteration dimension=1
//ff:what TestStoredDecompress — a zstd CLI frame decodes; bad header, bad block, bad checksum, trailing bytes and compressed blocks fail
package pack

import (
	"errors"
	"testing"
)

func TestStoredDecompress(t *testing.T) {
	cli := []byte{0x28, 0xb5, 0x2f, 0xfd, 0x24, 0x03, 0x19, 0x00, 0x00, 'a', 'b', 'c', 0x99, 0x09, 0x77, 0xad}
	if b, err := (Stored{}).Decompress(cli, nil); err != nil || string(b) != "abc" {
		t.Fatal(err)
	}
	noSum := []byte{0x28, 0xb5, 0x2f, 0xfd, 0x20, 0x03, 0x19, 0x00, 0x00, 'a', 'b', 'c'}
	if b, err := (Stored{}).Decompress(noSum, nil); err != nil || string(b) != "abc" {
		t.Fatal(err)
	}
	rep := []byte{0x28, 0xb5, 0x2f, 0xfd, 0x64, 0xe8, 0x02, 0x45, 0x00, 0x00, 0x08, 0x78, 0x01, 0x00, 0xe4, 0x2b, 0x20, 0x04, 0xe1, 0x00, 0xb7, 0x9c}
	if _, err := (Stored{}).Decompress(rep, nil); !errors.Is(err, ErrUnsupported) {
		t.Fatal(err)
	}
	bad := [][]byte{{1}, cli[:10], append(cli[:15:15], 0xae), append(append([]byte{}, cli...), 0)}
	for i, b := range bad {
		if _, err := (Stored{}).Decompress(b, nil); !errors.Is(err, ErrCorrupt) {
			t.Fatal(i, err)
		}
	}
}
