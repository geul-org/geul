//ff:func feature=core type=encoder control=sequence
//ff:what TestIssuerFromKey — a key gives a canonical 16-word Issuer, the same each time, different per key
package core

import (
	"slices"
	"testing"
)

func TestIssuerFromKey(t *testing.T) {
	a, err := IssuerFromKey(SigEd25519, make([]byte, 32))
	if err != nil || len(a) != 16 || !IssuerCanonical(a) {
		t.Fatalf("%04X %v", a, err)
	}
	again, _ := IssuerFromKey(SigEd25519, make([]byte, 32))
	other, _ := IssuerFromKey(SigEd25519, append(make([]byte, 31), 1))
	if !slices.Equal(a, again) || slices.Equal(a, other) {
		t.Fatal("not a function of the key")
	}
}
