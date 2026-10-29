//ff:func feature=pack type=encoder control=sequence
//ff:what TestStoredCompress — the raw-block frame of the input
package pack

import (
	"bytes"
	"testing"
)

func TestStoredCompress(t *testing.T) {
	f, err := Stored{}.Compress([]byte("x"))
	if err != nil || !bytes.Equal(f, zstdFrame([]byte("x"))) {
		t.Fatal(err)
	}
}
