//ff:func feature=pack type=encoder control=sequence
//ff:what TestStoredCompress — the raw-block frame of the input; a dictionary is ignored
package pack

import (
	"bytes"
	"testing"
)

func TestStoredCompress(t *testing.T) {
	f, err := Stored{}.Compress([]byte("x"), &Dict{ID: 9})
	if err != nil || !bytes.Equal(f, zstdFrame([]byte("x"))) {
		t.Fatal(err)
	}
}
