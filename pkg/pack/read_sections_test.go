//ff:func feature=pack type=decoder control=iteration dimension=1
//ff:what TestReadSections — bodies by kind; kinds ≥ 0x8000 allowed, duplicates, unknown required-range kinds, reserved bytes and bounds rejected
package pack

import (
	"encoding/binary"
	"errors"
	"testing"
)

func TestReadSections(t *testing.T) {
	entry := func(kind uint16, res byte, off, n uint64) []byte {
		b := binary.LittleEndian.AppendUint16(nil, kind)
		b = append(b, res, 0, 0, 0, 0, 0)
		b = binary.LittleEndian.AppendUint64(b, off)
		return binary.LittleEndian.AppendUint64(b, n)
	}
	whole := []byte("0123456789")
	secs, err := readSections(&cursor{b: append(entry(1, 0, 2, 3), entry(0x8001, 0, 0, 1)...)}, 2, whole)
	if err != nil || string(secs[1]) != "234" || len(secs) != 2 {
		t.Fatal(secs, err)
	}
	bad := [][]byte{
		append(entry(1, 0, 0, 1), entry(1, 0, 0, 1)...),
		entry(7, 0, 0, 1), entry(0, 0, 0, 1), entry(1, 1, 0, 1), entry(1, 0, 11, 0), entry(1, 0, 5, 6), entry(1, 0, 0, 1)[:20],
	}
	for i, b := range bad {
		n := uint32(len(b) / 24)
		if n == 0 {
			n = 1
		}
		if _, err := readSections(&cursor{b: b}, n, whole); !errors.Is(err, ErrCorrupt) {
			t.Fatal(i, err)
		}
	}
}
