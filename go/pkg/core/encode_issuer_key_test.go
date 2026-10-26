//ff:func feature=core type=encoder control=sequence
//ff:what TestEncodeIssuerKey — an Ed25519 key round-trips; another algorithm or key size fails
package core

import (
	"bytes"
	"testing"
)

func TestEncodeIssuerKey(t *testing.T) {
	pub := bytes.Repeat([]byte{7}, 32)
	r, err := EncodeIssuerKey(SigEd25519, pub)
	if err != nil || r.Codec != MetaIssuerKey || len(r.Payload) != 17 {
		t.Fatalf("%+v %v", r, err)
	}
	if alg, got, err := DecodeIssuerKey(r); err != nil || alg != SigEd25519 || !bytes.Equal(got, pub) {
		t.Fatalf("round trip: %v", err)
	}
	if _, err := EncodeIssuerKey(2, pub); err == nil {
		t.Fatal("algorithm 2: want error")
	}
	if _, err := EncodeIssuerKey(SigEd25519, pub[:30]); err == nil {
		t.Fatal("short key: want error")
	}
}
