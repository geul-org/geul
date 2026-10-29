//ff:func feature=core type=encoder control=sequence
//ff:what TestIssuerOfDigest — a digest with a high word set is the Issuer; one below 2^64 is refused
package core

import "testing"

func TestIssuerOfDigest(t *testing.T) {
	d := make([]byte, 32)
	d[0] = 0x80
	if ws, err := issuerOfDigest(d); err != nil || len(ws) != 16 || ws[0] != 0x8000 {
		t.Fatal(ws, err)
	}
	low := make([]byte, 32)
	low[31] = 1
	if _, err := issuerOfDigest(low); err == nil {
		t.Fatal("below 2^64: want error")
	}
}
