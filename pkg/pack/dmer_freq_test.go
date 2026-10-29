//ff:func feature=pack type=helper control=sequence
//ff:what TestDmerFreq — document frequency over samples
package pack

import (
	"encoding/binary"
	"testing"
)

func TestDmerFreq(t *testing.T) {
	f := dmerFreq([][]byte{[]byte("abcdefghabcdefgh"), []byte("abcdefgh"), []byte("zzzzzzzz")})
	if f[binary.LittleEndian.Uint64([]byte("abcdefgh"))] != 2 || f[binary.LittleEndian.Uint64([]byte("zzzzzzzz"))] != 1 {
		t.Fatal(f)
	}
}
