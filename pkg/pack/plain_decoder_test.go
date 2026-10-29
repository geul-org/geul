//ff:func feature=pack type=decoder control=sequence
//ff:what TestPlainDecoder — one shared decoder that reads raw-block frames
package pack

import "testing"

func TestPlainDecoder(t *testing.T) {
	d := plainDecoder()
	if d == nil || d != plainDecoder() {
		t.Fatal("not shared")
	}
	if b, err := d.DecodeAll(zstdFrame([]byte("abc")), nil); err != nil || string(b) != "abc" {
		t.Fatal(err)
	}
}
