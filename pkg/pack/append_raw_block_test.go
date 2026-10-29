//ff:func feature=pack type=encoder control=sequence
//ff:what TestAppendRawBlock — 3-byte header: size << 3, type 0, last bit
package pack

import (
	"bytes"
	"testing"
)

func TestAppendRawBlock(t *testing.T) {
	if b := appendRawBlock(nil, []byte("ab"), true); !bytes.Equal(b, []byte{0x11, 0, 0, 'a', 'b'}) {
		t.Fatalf("%x", b)
	}
	if b := appendRawBlock(nil, nil, false); !bytes.Equal(b, []byte{0, 0, 0}) {
		t.Fatalf("%x", b)
	}
}
