//ff:func feature=pack type=helper control=sequence
//ff:what TestZeroDmers — the d-mers of a segment drop to 0, others stay
package pack

import (
	"encoding/binary"
	"testing"
)

func TestZeroDmers(t *testing.T) {
	f := dmerFreq([][]byte{[]byte("abcdefghi"), []byte("zzzzzzzz")})
	zeroDmers(f, []byte("abcdefghi"))
	if f[binary.LittleEndian.Uint64([]byte("abcdefgh"))] != 0 || f[binary.LittleEndian.Uint64([]byte("bcdefghi"))] != 0 || f[binary.LittleEndian.Uint64([]byte("zzzzzzzz"))] != 1 {
		t.Fatal(f)
	}
}
