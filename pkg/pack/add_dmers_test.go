//ff:func feature=pack type=helper control=sequence
//ff:what TestAddDmers — each distinct 8-byte substring counts once per sample; short samples add nothing
package pack

import (
	"encoding/binary"
	"testing"
)

func TestAddDmers(t *testing.T) {
	freq, seen := map[uint64]uint32{}, map[uint64]bool{}
	addDmers(freq, seen, []byte("aaaaaaaaaaaa"))
	addDmers(freq, seen, []byte("aaaaaaaab"))
	addDmers(freq, seen, []byte("short"))
	a := binary.LittleEndian.Uint64([]byte("aaaaaaaa"))
	b := binary.LittleEndian.Uint64([]byte("aaaaaaab"))
	if len(freq) != 2 || freq[a] != 2 || freq[b] != 1 {
		t.Fatal(freq)
	}
}
