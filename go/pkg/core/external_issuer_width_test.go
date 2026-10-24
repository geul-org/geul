//ff:func feature=core type=helper control=iteration dimension=1
//ff:what TestExternalIssuerWidth — 1/2/4/16 for EXTERNAL-ISSUER-*, 0 for other Reference codecs
package core

import "testing"

func TestExternalIssuerWidth(t *testing.T) {
	want := map[uint8]int{RefExternalIssuer1Word: 1, RefExternalIssuer2Word: 2, RefExternalIssuer4Word: 4, RefExternalIssuer16Word: 16}
	for c := uint8(0); c < 16; c++ {
		if got := ExternalIssuerWidth(c); got != want[c] {
			t.Fatalf("%04b: %d want %d", c, got, want[c])
		}
	}
}
