//ff:func feature=pack type=helper control=iteration dimension=1
//ff:what TestXxh64 — reference values: empty, "a", "abc", and a 45-byte input checked against the zstd CLI checksum
package pack

import "testing"

func TestXxh64(t *testing.T) {
	b := make([]byte, 45)
	for i := range b {
		b[i] = byte(i)
	}
	if xxh64(nil) != 0xef46db3751d8e999 || xxh64([]byte("a")) != 0xd24ec4f1a98c6e5b || xxh64([]byte("abc")) != 0x44bc2cf5ad770999 || uint32(xxh64(b)) != 0x6409abdf {
		t.Fatal("xxh64")
	}
}
