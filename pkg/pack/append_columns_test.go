//ff:func feature=pack type=encoder control=sequence
//ff:what TestAppendColumns — column after column, varints
package pack

import (
	"bytes"
	"testing"
)

func TestAppendColumns(t *testing.T) {
	if b := appendColumns([]byte{9}, []uint64{1, 300}, []uint64{0}); !bytes.Equal(b, []byte{9, 1, 0xAC, 2, 0}) {
		t.Fatalf("%x", b)
	}
}
